import { useState, useEffect } from "react";
import { productService } from "../../services/productService.js";
import Input from "../../components/Input.jsx";
import Button from "../../components/Button.jsx";

const CategorySearch = ({ onSelectCategory, selectedCategoryId }) => {

  const [searchTerm, setSearchTerm] = useState("");
  const [categories, setCategories] = useState([]);
  const [displayCategories, setDisplayCategories] = useState([]);
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState(null);

  // Fetch categories on mount
  useEffect(() => {
    const fetchCategories = async () => {
      setIsLoading(true);
      try {
        const response = await productService.getCategories();
        const processedCategories = response.data.map(cat => ({
          ...cat,
          path: buildPath(cat, response.data)
        }));
        setCategories(response.data);
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

  // Build category path
  const buildPath = (category, allCategories) => {
    const path = [];
    let current = category;
    while (current) {
      path.unshift(current.name);
      current = allCategories.find(cat => cat.id === current.parent_id);
    }
    return path.join(" > ");
  };

  // Handle search
  const handleSearch = (e) => {
    const term = e.target.value.toLowerCase();
    setSearchTerm(term);
    
    const filtered = categories
      .filter(cat => 
        cat.name.toLowerCase().includes(term) ||
        cat.path.toLowerCase().includes(term) ||
        cat.slug.toLowerCase().includes(term)
      )
      .map(cat => ({ ...cat, path: buildPath(cat, categories) }));
    
    setDisplayCategories(term ? filtered : categories.map(cat => ({
      ...cat,
      path: buildPath(cat, categories)
    })));
  };

  const clearSearch = () => {
    setSearchTerm("");
    setDisplayCategories(categories.map(cat => ({
      ...cat,
      path: buildPath(cat, categories)
    })));
  };

  return (
    <div className="bg-white p-4 rounded-lg shadow">
      <h2 className="text-lg font-bold mb-3">Category Khojein</h2>
      
      <div className="flex gap-2 mb-3">
        <Input
          placeholder="Name, path ya slug se khojein..."
          value={searchTerm}
          onChange={handleSearch}
          className="flex-1"
        />
        <Button onClick={clearSearch} variant="secondary">Clear</Button>
      </div>

      {isLoading && <p>Loading...</p>}
      {error && <p className="text-red-500">{error}</p>}
      
      {!isLoading && !error && (
        <div className="overflow-auto max-h-72 border rounded">
          {displayCategories.length > 0 ? (
            <table className="w-full">
              <thead className="bg-gray-100 sticky top-0">
                <tr className="text-left text-sm text-gray-600">
                  <th className="p-2">Name</th>
                  <th className="p-2">Path</th>
                  <th className="p-2">Status</th>
                  <th className="p-2">Action</th>
                </tr>
              </thead>
              <tbody>
                {displayCategories.map(category => 
                  selectedCategoryId !== category.id && (
                    <tr key={category.id} className="border-t hover:bg-gray-50">
                      <td className="p-2">
                        <div>{category.name}</div>
                        <div className="text-xs text-gray-500">{category.slug}</div>
                      </td>
                      <td className="p-2 text-sm">{category.path}</td>
                      <td className="p-2">
                        <span className={`px-2 py-1 text-xs rounded-full ${
                          category.status === "active" 
                            ? "bg-green-100 text-green-800" 
                            : "bg-red-100 text-red-800"
                        }`}>
                          {category.status === "active" ? "Active" : "Inactive"}
                        </span>
                      </td>
                      <td className="p-2">
                        <Button 
                          size="sm" 
                          onClick={() => onSelectCategory(category)}
                        >
                          Select
                        </Button>
                      </td>
                    </tr>
                  )
                )}
              </tbody>
            </table>
          ) : (
            <p className="p-4 text-center text-gray-500">Koi category nahi mili</p>
          )}
        </div>
      )}
    </div>
  );
};

export default CategorySearch;