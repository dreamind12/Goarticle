import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import ReactPaginate from 'react-paginate';

const PreviewPost = () => {
  const [postData, setPostData] = useState([]);
  const [pageNumber, setPageNumber] = useState(0);
  const postsPerPage = 3; 

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetch(`http://localhost:8080/article/page/${postsPerPage}/${pageNumber}`);
        const data = await response.json();
        setPostData(data.data);
      } catch (error) {
        console.error('Error fetching data:', error);
      }
    };

    fetchData();
  }, [pageNumber]);

  const handlePageClick = (selected) => {
    setPageNumber(selected.selected);
  };

   return (
    <div className="container mx-auto">
      <nav className="bg-gray-800 text-white py-4 px-6 flex items-center justify-between">
        <Link to="/"> <h2 className="text-2xl font-bold mb-4">Preview Post</h2></Link>
      </nav>

      
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4 p-4">
        {postData.map((post) => (
          <div key={post.ID} className="bg-white rounded-md shadow-md p-4">
            <h3 className="text-xl font-bold mb-2">{post.Title}</h3>
            <p className="mb-2">{post.Content}</p>
            <div className="flex justify-between">
              <p className="text-sm">{post.Category}</p>
              <p className={`text-sm ${post.Status === 'Publish' ? 'text-green-500' : 'text-gray-500'}`}>
                {post.Status}
              </p>
            </div>
          </div>
        ))}
      </div>

      
      <ReactPaginate
        pageCount={Math.ceil(postData.length / postsPerPage)}
        pageRangeDisplayed={3}
        marginPagesDisplayed={1}
        onPageChange={handlePageClick}
        containerClassName={'pagination'}
        activeClassName={'active'}
      />
    </div>
  );
};

export default PreviewPost;