import { useEffect, useState, useActionState, useRef } from "react";
import { useNavigate } from "react-router-dom";
import { ToastContainer, toast } from "react-toastify";
import { productService } from "../../services/productService";

import Button from "../../components/Button";
import Input from "../../components/Input";
import Label from "../../components/Label";
import Select from "../../components/Select";
import BackButton from "../../components/BackButton";
import CategorySearch from "./CategorySearch";

import RichTextEditor from "../../components/AdvancedRichTextEditor";

const AddCategory = () => {
  const navigate = useNavigate();
 


  const [parentCategories, setParentCategories] = useState([]);
  const [categoryData, setCategoryData] = useState({
    name: "",
    slug: "",
    description: "",
    parent_id: "",
    status: "active"
  });
  const [isLoading, setIsLoading] = useState(true);
  const [selectedCategory, setSelectedCategory] = useState(null);
  const [showCategorySearch, setShowCategorySearch] = useState(false);
  const [autoGenerateSlug, setAutoGenerateSlug] = useState(true);
  const formRef = useRef(null);

  const handleDescriptionChange = (newValue) => {
    setCategoryData((prev) => ({ ...prev, description: newValue }));
  };

  // Function to transform flat category list into hierarchical structure
  const buildCategoryTree = (categories, parentId = "") => {
    return categories
      .filter(cat => String(cat.parent_id || "") === String(parentId))
      .map(cat => ({
        ...cat,
        children: buildCategoryTree(categories, cat.id)
      }));
  };

  // Function to create options for Select with proper indentation
  const createSelectOptions = (categoryTree, level = 0, currentCategoryId = null, parentPath = "") => {
    let options = [];
    
    categoryTree.forEach(category => {
      // Skip the current category and its descendants to prevent circular references
      if (!currentCategoryId || category.id !== parseInt(currentCategoryId)) {
        const currentPath = parentPath ? `${parentPath} > ${category.name}` : category.name;
        
        options.push({
          value: String(category.id),
          name: "  ".repeat(level) + (level > 0 ? "└ " : "") + category.name,
          path: currentPath
        });
        
        if (category.children && category.children.length > 0) {
          options = [...options, ...createSelectOptions(category.children, level + 1, currentCategoryId, currentPath)];
        }
      }
    });
    
    return options;
  };
  
  useEffect(() => {
    // Fetch parent categories
    const fetchData = async () => {
      setIsLoading(true);
      try {
        const parentsResponse = await productService.getCategories();
        
        // Build hierarchical tree for better visualization
        const allCategories = parentsResponse.data;
        const categoryTree = buildCategoryTree(allCategories);
        const hierarchicalOptions = createSelectOptions(categoryTree);
        
        setParentCategories([{ value: "", name: "None", path: "" }, ...hierarchicalOptions]);
        setIsLoading(false);
      } catch (err) {
        toast.error("Failed to load categories");
        console.error(err);
        setIsLoading(false);
      }
    };
    
    fetchData();
  }, [url, categoriesUrl]);

  const handleSelectCategory = (category) => {
    setSelectedCategory(category);
    setShowCategorySearch(false);
    
    // Update parent_id with the selected category's id
    setCategoryData(prev => ({
      ...prev,
      parent_id: category.id
    }));
    
    toast.info(`Selected parent category: ${category.name}`);
  };

  // Handle name change to generate slug
  const handleNameChange = (e) => {
    const nameValue = e.target.value;
    
    // Update categoryData with the new name
    setCategoryData(prev => ({
      ...prev,
      name: nameValue
    }));
    
    // Handle auto slug generation
    if (autoGenerateSlug) {
      // Generate slug from name
      const generatedSlug = nameValue
        .toLowerCase()
        .replace(/[^\w\s-]/g, '') 
        .replace(/\s+/g, '-')
        .replace(/-+/g, '-');
      
      // Update slug in state
      setCategoryData(prev => ({
        ...prev,
        slug: generatedSlug
      }));
    }
  };

  // Handle slug change
  const handleSlugChange = (e) => {
    if (!autoGenerateSlug) {
      setCategoryData(prev => ({
        ...prev,
        slug: e.target.value
      }));
    }
  };

  // Toggle auto slug generation
  const toggleAutoSlug = () => {
    setAutoGenerateSlug(!autoGenerateSlug);
  };

  const [error, handleSubmit, isPending] = useActionState(async (previousState, formData) => {
    // Create category data object from form data
    const categoryDataObj = {
      name: formData.get("name"),
      slug: formData.get("slug"),
      description: categoryData.description,
      parent_id: formData.get("parent_id") || null,
      status: formData.get("status")
    };

    try {
      const response = await productService.createCategory(categoryDataObj);
      toast.success("Category created successfully!");
      
      // Redirect after successful creation
      setTimeout(() => navigate("/categories"), 2000);
      return { error: null };
    } catch (err) {
      toast.error(err.message || "Failed to create category");
      return { error: err.message || "An unexpected error occurred" };
    }
  }, { error: null });

  if (isLoading) {
    return (
      <div className="flex justify-center items-center h-64">
        <div className="animate-spin h-12 w-12 border-4 border-blue-500 border-t-transparent rounded-full"></div>
      </div>
    );
  }

  return (
    <>
      <ToastContainer position="top-right" autoClose={3000} />
      <div className="bg-white shadow-lg rounded-lg p-6">
        <div className="flex justify-between items-center mb-6">
          <BackButton>Back to Categories</BackButton>
          <h1 className="text-2xl font-bold">➕ Add New Category</h1>
          <Button 
            onClick={() => setShowCategorySearch(!showCategorySearch)}
            className="bg-indigo-600 hover:bg-indigo-700"
          >
            {showCategorySearch ? "Hide Search" : "Find Parent"}
          </Button>
        </div>
        
        {showCategorySearch && (
          <CategorySearch 
            onSelectCategory={handleSelectCategory} 
          />
        )}
        

        <form ref={formRef} action={handleSubmit} className="mt-6 space-y-6">
          <div>
            <Label htmlFor="name">Category Name:</Label>
            <Input 
              type="text" 
              id="name" 
              name="name" 
              value={categoryData.name} 
              placeholder="Enter category name"
              onChange={handleNameChange}
            />
          </div>

          <div>
            <div className="flex justify-between items-center">
              <Label htmlFor="slug">Slug:</Label>
              <div className="flex items-center gap-2">
                <input 
                  type="checkbox" 
                  id="autoSlug" 
                  checked={autoGenerateSlug} 
                  onChange={toggleAutoSlug}
                  className="h-4 w-4 text-blue-600"
                />
                <label htmlFor="autoSlug" className="text-sm text-gray-600">Auto-generate from name</label>
              </div>
            </div>
            <Input 
              type="text" 
              id="slug" 
              name="slug" 
              value={categoryData.slug} 
              placeholder="enter-category-slug"
              readOnly={autoGenerateSlug}
              className={autoGenerateSlug ? "bg-gray-100" : ""}
              onChange={handleSlugChange}
            />
          </div>

          <div>
            <Label htmlFor="parent_id">Parent Category:</Label>
            
            <Select 
              id="parent_id"
              name="parent_id" 
              data={parentCategories} 
              value={categoryData.parent_id || "0"} 
              onChange={(event) => {
                setCategoryData(prev => ({
                  ...prev,
                  parent_id: event.target.value
                }));  
              }}
            />
            
            {/* Show path information for the selected parent */}
            {categoryData.parent_id && (
              <div className="mt-1 text-sm text-gray-500">
                Full path: {parentCategories.find(p => p.value === String(categoryData.parent_id))?.path || ""}
              </div>
            )}
          </div>

          <div>
            <Label htmlFor="description">Description:</Label>
            <RichTextEditor
              id="description"
              name="description"
              value={categoryData.description} 
              onChange={handleDescriptionChange}
              rows={4}
              maxLength={10000}
              placeholder="Write a detailed category description..."
            />
          </div>

          <div>
            <Label htmlFor="image">Image:</Label>
            <input
              type="file"
              id="image"
              name="image"
              accept="image/*"
              className="w-full border border-gray-400 rounded px-4 py-2 text-lg focus:ring-2 focus:ring-green-500 transition"
            />
            <p className="text-sm text-gray-500 mt-1">
              Recommended: Square image (500×500px), Max: 2MB
            </p>
          </div>

          <div>
            <Label htmlFor="status">Status:</Label>
            <Select 
              id="status"
              name="status" 
              data={[
                { value: "active", name: "Active" },
                { value: "inactive", name: "Inactive" },
              ]} 
              value={categoryData.status} 
              onChange={(event) => {
                setCategoryData(prev => ({
                  ...prev,
                  status: event.target.value
                }));  
              }}
            />
          </div>

          <div className="flex justify-between pt-4">
            <Button 
              type="button" 
              onClick={() => navigate(-1)}
              className="bg-gray-600 hover:bg-gray-700"
            >
              ❌ Cancel
            </Button>
            <Button 
              type="submit" 
              disabled={isPending}
              className="bg-green-600 hover:bg-green-700"
            >
              {isPending ? (
                <span className="flex items-center">
                  <span className="animate-spin h-4 w-4 mr-2 border-2 border-white border-t-transparent rounded-full"></span>
                  Creating...
                </span>
              ) : (
                "✅ Create Category"
              )}
            </Button>
          </div>
        </form>
      </div>
    </>
  );
};

export default AddCategory;