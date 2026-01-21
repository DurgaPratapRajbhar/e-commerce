import { useState, useEffect, useRef } from "react";
import { useNavigate } from "react-router-dom";
import { PencilSquareIcon, TrashIcon, PlusCircleIcon } from "@heroicons/react/24/solid";
import { api } from "../../lib/api";

const CategoryList = () => {
  
  const env=import.meta.env;
  const url=env.VITE_PRODUCT_SERVICE;
  const categoriesUrl=env.VITE_PRODUCT_CATEGORIES;
  
  const [categories, setCategories] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const isMounted = useRef(false);
  const navigate = useNavigate();

  useEffect(() => {
    if (!isMounted.current) {
      fetchCategories();
      isMounted.current = true;
    }
  }, []);

  const fetchCategories = async () => {
    try {
      const response = await api.get(`${url}${categoriesUrl}`);
      setCategories(buildCategoryTree(response.data.data));
    } catch (err) {
      setError("Failed to load categories.");
      console.error("Error:", err);
    } finally {
      setLoading(false);
    }
  };

  // Convert flat array to tree structure and ensure `children` exists
  const buildCategoryTree = (data) => {
    let categoryMap = {};
    let tree = [];

    data.forEach((category) => {
      categoryMap[category.id] = { ...category, children: category.children || [], expanded: false };
    });

    data.forEach((category) => {
      if (category.parent_id && categoryMap[category.parent_id]) {
        categoryMap[category.parent_id].children.push(categoryMap[category.id]);
      } else {
        tree.push(categoryMap[category.id]);
      }
    });

    return tree;
  };

  const deleteCategory = async (id) => {
    if (!confirm(`Do you really want to delete category ${id}?`)) return;
    try {
      await api.delete(`${url}${categoriesUrl}/${id}`);
      setCategories((prevCategories) => buildCategoryTree(prevCategories.filter((category) => category.id !== id)));
    } catch (err) {
      console.error("Delete Error:", err);
    }
  };

  return (
    <div className="bg-white shadow-lg rounded-lg p-6">
      <div className="flex justify-between items-center mb-4">
        <h2 className="text-2xl font-semibold text-gray-800">📂 Category List</h2>
        <button
          className="bg-green-500 text-white px-4 py-2 rounded-lg flex items-center gap-2 hover:bg-green-600 transition"
          onClick={() => navigate("/categories/add")}
        >
          <PlusCircleIcon className="h-5 w-5" /> Add Category
        </button>
      </div>

      {loading && <p className="text-center text-gray-500">Loading categories...</p>}
      {error && <p className="text-center text-red-500">{error}</p>}

      {!loading && !error && (
        <CategoryTree categories={categories} setCategories={setCategories} deleteCategory={deleteCategory} navigate={navigate} />
      )}
    </div>
  );
};

// Recursive Category Tree Component with Expand/Collapse
const CategoryTree = ({ categories, setCategories, deleteCategory, navigate, level = 0 }) => {
  const toggleExpand = (event, id) => {
    event.preventDefault(); // ✅ Prevents event bubbling issues

    const updateCategories = (categoryList) =>
      categoryList.map((category) => ({
        ...category,
        expanded: category.id === id ? !category.expanded : false, // ✅ Hide others when one expands
        children: category.children.map((child) => ({
          ...child,
          expanded: child.id === id ? !child.expanded : child.expanded,
        })),
      }));

    setCategories((prevCategories) => updateCategories(prevCategories));
  };

  return (
    <ul className={`pl-${level * 4}`}>
      {categories.map((category) => (
        <li key={category.id} className="border-l pl-4 mb-2">
          <div className="flex items-center gap-4">
            {/* Expand/Collapse Button */}
            {Array.isArray(category.children) && category.children.length > 0 && (
              <button
                className="text-gray-700 font-bold"
                onClick={(event) => toggleExpand(event, category.id)}
              >
                {category.expanded ? "➖" : "➕"}
              </button>
            )}

            <span className="font-semibold">{category.name}</span>
            <span className="text-gray-500 text-sm">({category.slug})</span>

            <button className="text-blue-500 hover:text-blue-700 transition" onClick={() => navigate(`/categories/edit/${category.id}`)}>
              <PencilSquareIcon className="h-5 w-5" />
            </button>
            <button className="text-red-500 hover:text-red-700 transition" onClick={() => deleteCategory(category.id)}>
              <TrashIcon className="h-5 w-5" />
            </button>
          </div>

          {/* Subcategories - Expand/Collapse Feature */}
          {category.expanded && Array.isArray(category.children) && category.children.length > 0 && (
            <CategoryTree categories={category.children} setCategories={setCategories} deleteCategory={deleteCategory} navigate={navigate} level={level + 1} />
          )}
        </li>
      ))}
    </ul>
  );
};

export default CategoryList;
