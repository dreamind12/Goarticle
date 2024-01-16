import React from 'react';
import AllPosts from './pages/allPost';
import { BrowserRouter, Route, Routes, } from 'react-router-dom';
import CreatePost from './pages/createPost';
import UpdatePost from './pages/updatePost';
import PreviewPost from './pages/previewPost';

function App() {
  return (
     <BrowserRouter>
      <Routes>
        <Route path="/" element={<AllPosts />} />
        <Route path="/createPost" element={<CreatePost />} />
        <Route path="/updatePost/:postId" element={<UpdatePost />} />
        <Route path="/previewPost" element={<PreviewPost />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
