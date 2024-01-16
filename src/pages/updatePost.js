import React, { useState, useEffect } from 'react';
import { Link, useParams, useNavigate } from 'react-router-dom';
import axios from 'axios';

const UpdatePost = () => {
  const { postId } = useParams();
  const [postData, setPostData] = useState({
    title: '',
    content: '',
    category: '',
    status: 'Draft',
  });
  const navigate = useNavigate();

  useEffect(() => {
    const fetchPostData = async () => {
      try {
        const response = await axios.get(`http://localhost:8080/article/${postId}`);
        const post = response.data?.data; 
        if (post) {
          setPostData({
            title: post.Title,
            content: post.Content,
            category: post.Category,
            status: post.Status,
          });
        } else {
          console.error('Empty post data.');
        }
      } catch (error) {
        console.error('Error fetching post data:', error);
      }
    };
    

    fetchPostData();
  }, [postId]);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setPostData({ ...postData, [name]: value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      const response = await axios.put(`http://localhost:8080/article/${postId}`, postData);
      console.log('Post updated successfully:', response.data);

      navigate('/');
    } catch (error) {
      console.error('Error updating post:', error);
    }
  };

  return (
    <div className="container mx-auto">
      <nav className="bg-gray-800 text-white py-4 px-6 flex items-center justify-between">
        <Link to="/">
          <h2 className="text-2xl font-bold mb-4">Update Post</h2>
        </Link>
      </nav>

      <form onSubmit={handleSubmit} className="max-w-md mx-auto mt-8">
        <div className="mb-4">
          <label htmlFor="title" className="block text-sm font-semibold mb-1">
            Title:
          </label>
          <input
            type="text"
            id="title"
            name="title"
            value={postData.title}
            onChange={handleChange}
            className="w-full p-2 border rounded"
            required
          />
        </div>
        <div className="mb-4">
          <label htmlFor="content" className="block text-sm font-semibold mb-1">
            Content:
          </label>
          <textarea
            id="content"
            name="content"
            value={postData.content}
            onChange={handleChange}
            className="w-full p-2 border rounded"
            required
          ></textarea>
        </div>
        <div className="mb-4">
          <label htmlFor="category" className="block text-sm font-semibold mb-1">
            Category:
          </label>
          <input
            type="text"
            id="category"
            name="category"
            value={postData.category}
            onChange={handleChange}
            className="w-full p-2 border rounded"
            required
          />
        </div>
        <div className="mb-4">
          <label htmlFor="status" className="block text-sm font-semibold mb-1">
            Status:
          </label>
          <select
            id="status"
            name="status"
            value={postData.status}
            onChange={handleChange}
            className="w-full p-2 border rounded"
          >
            <option value="Draft">Draft</option>
            <option value="Publish">Publish</option>
          </select>
        </div>
        <button type="submit" className="bg-blue-500 text-white p-2 rounded">
          Update Post
        </button>
      </form>
    </div>
  );
};

export default UpdatePost;
