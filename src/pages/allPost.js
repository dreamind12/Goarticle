import React, { useState, useEffect } from 'react';
import { Link, useNavigate } from 'react-router-dom';

const AllPosts = () => {
  const [currentTab, setCurrentTab] = useState('published');
  const [postData, setPostData] = useState([]);
  const [selectedPost, setSelectedPost] = useState(null);
  const navigate = useNavigate();

  const tabs = ['published', 'drafts', 'trashed'];

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetch('http://localhost:8080/article/getAll');
        const data = await response.json();
        setPostData(data.data);
      } catch (error) {
        console.error('Error fetching data:', error);
      }
    };

    fetchData();
  }, []);

  const handleTabChange = (tab) => {
    setCurrentTab(tab);
    setSelectedPost(null); 
  };

const handleEditClick = (post) => {
  setSelectedPost(post);
  navigate(`/updatePost/${post.ID}`);
};

const handleTrashClick = async (post) => {
  try {
    await fetch(`http://localhost:8080/article/${post.ID}`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        status: 'Thrash',
      }),
    });

    setPostData((prevData) => {
      return prevData.map((prevPost) =>
        prevPost.ID === post.ID ? { ...prevPost, Status: 'Thrash' } : prevPost
      );
    });
  } catch (error) {
    console.error('Error trashing post:', error);
  }
};


  const getFilteredData = () => {
    switch (currentTab) {
      case 'published':
        return postData.filter((post) => post.Status === 'Publish');
      case 'drafts':
        return postData.filter((post) => post.Status === 'Draft');
      case 'trashed':
        return postData.filter((post) => post.Status === 'Thrash');
      default:
        return [];
    }
  };

  const renderTable = () => {
    const filteredData = getFilteredData();

    return (
      <div className="p-4 m-4 bg-white rounded-md shadow-md">
        <table className="table-auto w-full">
          <thead>
            <tr>
              <th className="px-4 py-2">Title</th>
              <th className="px-4 py-2">Category</th>
              <th className="px-4 py-2">Action</th>
            </tr>
          </thead>
          <tbody>
            {filteredData.map((post) => (
              <tr key={post.ID}>
                <td className="border px-4 py-2">{post.Title}</td>
                <td className="border px-4 py-2">{post.Category}</td>
                <td className="border px-4 py-2">
                  <button className="mr-2 bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded" onClick={() => handleEditClick(post)}>
                     Edit
                  </button>
                  <button className='bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded'  onClick={() => handleTrashClick(post)}>
                     Trash
                  </button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    );
  };

  return (
    <div className="bg-gray-800 p-4 mb-4 rounded-md shadow-md text-center">
      <div className="flex justify-center rounded-md mt-3">
        <div className="flex">
          {tabs.map((tab) => (
            <button
              key={tab}
              className={`mr-4 px-4 py-2 text-white ${
                currentTab === tab ? 'bg-blue-500' : 'bg-gray-600'
              }`}
              onClick={() => handleTabChange(tab)}
            >
              {tab.charAt(0).toUpperCase() + tab.slice(1)}
            </button>
          ))}
        </div>
        <Link to="/previewPost">
          <button className="justify-end px-4 py-2 mr-3 text-white bg-violet-500">Preview Post</button>
        </Link>
        <Link to="/createPost">
          <button className="px-4 py-2 text-white bg-green-500">Add New</button>
        </Link>
      </div>
      <div className="container mx-auto mt-8">{renderTable()}</div>
    </div>
  );
};

export default AllPosts;
