import React from "react";

const Pagination = ({ currentPage, totalPages, onPageChange }) => {
  return (
    <div className="mt-6 flex items-center justify-between">
      <button
        className={`px-3 py-1 rounded-lg ${currentPage === 1 ? 'bg-gray-200 cursor-not-allowed' : 'bg-indigo-600 text-white hover:bg-indigo-700'}`}
        onClick={() => onPageChange(currentPage - 1)}
        disabled={currentPage === 1}
      >
        Previous
      </button>
      <span className="px-3 py-1 text-sm text-gray-600">
        Page {currentPage} of {totalPages}
      </span>
      <button
        className={`px-3 py-1 rounded-lg ${currentPage === totalPages ? 'bg-gray-200 cursor-not-allowed' : 'bg-indigo-600 text-white hover:bg-indigo-700'}`}
        onClick={() => onPageChange(currentPage + 1)}
        disabled={currentPage === totalPages}
      >
        Next
      </button>
    </div>
  );
};

export default Pagination;
