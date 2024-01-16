import React, { useState } from 'react';
import axios from 'axios';
import { Link, useNavigate} from 'react-router-dom';

const CreatePost = () => {
  const [postData, setPostData] = useState({
    title: '',
    content: '',
    category: '',
    status: 'Draft', // Default status
  });
  const navigate = useNavigate();

  const handleChange = (e) => {
    const { name, value } = e.target;
    setPostData({ ...postData, [name]: value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      const response = await axios.post('http://localhost:8080/article/', postData);
      console.log('Post created successfully:', response.data);
      navigate('/');
      setPostData({
        title: '',
        content: '',
        category: '',
        status: 'Draft',
      });
    } catch (error) {
      console.error('Error creating post:', error);
    }
  };

  return (
    <div className="container mx-auto">
        <nav className="bg-gray-800 text-white py-4 px-6 flex items-center justify-between">
        <Link to="/"> <h2 className="text-2xl font-bold mb-4">Create Post</h2></Link>
        
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
          Create Post
        </button>
      </form>
    </div>
  );
};

export default CreatePost;
