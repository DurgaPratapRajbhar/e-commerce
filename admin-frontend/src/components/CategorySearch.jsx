import { useState, useEffect } from "react";
import { api } from "../lib/api";

const CategorySearch = ({ onSelectCategory, selectedCategoryId, setShowCategoryModal }) => {
  const { VITE_PRODUCT_SERVICE: url, VITE_PRODUCT_CATEGORIES: categoriesUrl } = import.meta.env;

  const [searchTerm, setSearchTerm] = useState("");
  const [categories, setCategories] = useState([]);
  const [displayCategories, setDisplayCategories] = useState([]);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchCategories = async () => {
      setIsLoading(true);
      try {
        const response = await api.get(`${url}${categoriesUrl}`);
        const data = response.data;
        const processedCategories = data.map(cat => ({
          ...cat,
          path: buildPath(cat, data),
        }));
        setCategories(data);
        setDisplayCategories(processedCategories);
      } catch (err) {
        setError("Categories load karne mein problem hui");
        console.error(err);
      } finally {
        setIsLoading(false);
      }
    };
    fetchCategories();
  }, []);

  const buildPath = (category, allCategories) => {
    const path = [];
    let current = category;
    while (current) {
      path.unshift(current.name);
      current = allCategories.find(cat => cat.id === current.parent_id);
    }
    return path.join(" > ");
  };

  const handleSearch = (e) => {
    const term = e.target.value.toLowerCase();
    setSearchTerm(term);

    const filtered = categories
      .filter(cat =>
        cat.name.toLowerCase().includes(term) ||
        buildPath(cat, categories).toLowerCase().includes(term) ||
        cat.slug.toLowerCase().includes(term)
      )
      .map(cat => ({
        ...cat,
        path: buildPath(cat, categories),
      }));

    setDisplayCategories(filtered);
  };

  const clearSearch = () => {
    setSearchTerm("");
    const refreshed = categories.map(cat => ({
      ...cat,
      path: buildPath(cat, categories),
    }));
    setDisplayCategories(refreshed);
  };

  const handleSelectCategory = (category) => {
    onSelectCategory(category);
    setShowCategoryModal(false);
  };

  return (
    <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
      <div className="bg-white rounded-lg p-6 w-3/4 max-w-3xl max-h-[80vh] overflow-auto">
        <div className="flex justify-between items-center mb-4">
          <h3 className="text-xl font-semibold">Select Category</h3>
          <button
            onClick={() => setShowCategoryModal(false)}
            className="text-gray-500 hover:text-gray-700"
          >
            ✕
          </button>
        </div>

        <input
          type="text"
          placeholder="Search categories..."
          className="w-full p-2 border rounded mb-4"
          value={searchTerm}
          onChange={handleSearch}
        />

        {error && <p className="text-red-500 mb-2">{error}</p>}
        {isLoading ? (
          <p>Loading...</p>
        ) : (
          <div className="overflow-auto max-h-96">
            <table className="w-full">
              <thead className="bg-gray-100 sticky top-0">
                <tr className="text-left">
                  <th className="p-2">Name</th>
                  <th className="p-2">Path</th>
                  <th className="p-2">Action</th>
                </tr>
              </thead>
              <tbody>
                {displayCategories.length > 0 ? (
                  displayCategories.map(category =>
                    selectedCategoryId !== category.id ? (
                      <tr key={category.id} className="border-t hover:bg-gray-50">
                        <td className="p-2">{category.name}</td>
                        <td className="p-2 text-sm text-gray-600">{category.path}</td>
                        <td className="p-2">
                          <button
                            onClick={() => handleSelectCategory(category)}
                            className="bg-blue-500 hover:bg-blue-600 text-white px-3 py-1 rounded text-sm"
                          >
                            Select
                          </button>
                        </td>
                      </tr>
                    ) : null
                  )
                ) : (
                  <tr>
                    <td colSpan="3" className="text-center p-4 text-gray-500">
                      No categories found.
                    </td>
                  </tr>
                )}
              </tbody>
            </table>
          </div>
        )}
      </div>
    </div>
  );
};

export default CategorySearch;
