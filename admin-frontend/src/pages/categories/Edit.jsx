import { useEffect, useState, useActionState, useRef } from "react";
import { useParams, useNavigate } from "react-router-dom";
import { ToastContainer, toast } from "react-toastify";
import { api } from "../../lib/api";

import Button from "../../components/Button";
import Input from "../../components/Input";
import Label from "../../components/Label";
import Select from "../../components/Select";
// import Textarea from "../../components/Textarea";
import BackButton from "../../components/BackButton";
import CategorySearch from "./CategorySearch";

import RichTextEditor from "../../components/AdvancedRichTextEditor";
 

const Edit = () => {
  const { id } = useParams();
  const navigate = useNavigate();
 
  const env = import.meta.env;
  const url = env.VITE_PRODUCT_SERVICE;
  const categoriesUrl = env.VITE_PRODUCT_CATEGORIES;

  const [parentCategories, setParentCategories] = useState([]);
  const [categoryData, setCategoryData] = useState({});
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
  const createSelectOptions = (categoryTree, level = 0, currentCategoryId, parentPath = "") => {
    let options = [];
    
    categoryTree.forEach(category => {
      // Skip the current category and its descendants to prevent circular references
      if (category.id !== parseInt(currentCategoryId)) {
        const currentPath = parentPath ? `${parentPath} > ${category.name}` : category.name;
        
        options.push({
          value: String(category.id),
          name: "  ".repeat(level) + (level > 0 ? "└ " : "") + category.name,
          path: currentPath
        });
        
        if (category.children && category.children.length > 0) {
          // Only include children if this category is not the one being edited
          options = [...options, ...createSelectOptions(category.children, level + 1, currentCategoryId, currentPath)];
        }
      }
    });
    
    return options;
  };
  
  useEffect(() => {
    // Fetch both category data and parent categories concurrently
    const fetchData = async () => {
      setIsLoading(true);
      try {
        const [categoryResponse, parentsResponse] = await Promise.all([
          api.get(`${url}${categoriesUrl}/${id}`),
          api.get(`${url}${categoriesUrl}`)
        ]);
        
        setCategoryData(categoryResponse.data.data);
        
        // Check for potential circular dependencies
        const allCategories = parentsResponse.data.data;
        
        // Build hierarchical tree for better visualization
        const categoryTree = buildCategoryTree(allCategories);
        const hierarchicalOptions = createSelectOptions(categoryTree, 0, parseInt(id));
        
        setParentCategories([{ value: "", name: "None", path: "" }, ...hierarchicalOptions]);
        setIsLoading(false);
      } catch (err) {
        toast.error("Failed to load category data");
        console.error(err);
        setIsLoading(false);
      }
    };
    
    fetchData();
  }, [id, url, categoriesUrl]);

  const handleSelectCategory = (category) => {
    setSelectedCategory(category);
    setShowCategorySearch(false);
    
    // Populate form with selected category data
    if (formRef.current) {
      const form = formRef.current;
      form.elements.name.value = category.name || "";
      form.elements.slug.value = category.slug || "";
      form.elements.description.value = category.description || "";
      form.elements.status.value = category.status || "active";
      
      if (form.elements.parent_id) {
        form.elements.parent_id.value = category.parent_id || "";
      }
    }
    
    // Update the categoryData state
    setCategoryData(category);
    
    toast.info(`Selected category: ${category.name}`);
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
    if (autoGenerateSlug && formRef.current) {
      const slugInput = formRef.current.elements.slug;
      
      // Generate slug from name
      const generatedSlug = nameValue
        .toLowerCase()
        .replace(/[^\w\s-]/g, '') 
        .replace(/\s+/g, '-')
        .replace(/-+/g, '-');
      
      slugInput.value = generatedSlug;
      
      // Also update slug in state
      setCategoryData(prev => ({
        ...prev,
        slug: generatedSlug
      }));
    }
  };

  // Toggle auto slug generation
  const toggleAutoSlug = () => {
    setAutoGenerateSlug(!autoGenerateSlug);
  };

  const [error, handleSubmit, isPending] = useActionState(async (previousState, formData) => {
    // Validate form data
    let errors = {};
    errors.name = formValidation(formData.get("name"), "Name");
    errors.slug = validateSlug(formData.get("slug"));
    //errors.description = formValidation(formData.get("description"), "Description", 10, 5000);
    errors.status = formValidation(formData.get("status"), "Status");

    // Validate parent_id to prevent circular dependencies
    const parentId = formData.get("parent_id");
    formData.append("description", categoryData.description);
 
    
    // Server should handle this check too, but we do it client-side as well
    if (parentId && parentId === id) {
      errors.parent_id = "Category cannot be its own parent";
    }

    // Validate image
    const categoryImage = formData.get("image");
    // if (categoryImage && categoryImage.size > 0) {
    //   const imageError = await validateImage(categoryImage, "Category Image", {
    //     maxSize: 2 * 1024 * 1024, // 2MB
    //     minWidth: 500,
    //     minHeight: 500,
    //     square: true,
    //   });
      
    //   if (imageError) {
    //     errors.image = imageError;
    //   }
    // }

    // Filter out empty error messages
    const filteredErrors = Object.fromEntries(
      Object.entries(errors).filter(([_, value]) => value)  
    );

    if (Object.keys(filteredErrors).length > 0) {
      toast.error("Please fix the validation errors");
      return filteredErrors;
    }

    try {
      // If no new image was selected, we need to keep the old one
      // if (!categoryImage || categoryImage.size === 0) {
      //   formData.delete("image");
      // }
      
      const response = await api.put(`${url}${categoriesUrl}/${id}`, formData, {
        headers: {
          'Content-Type': 'multipart/form-data',
        },
      });
      toast.success("Category updated successfully!");
      
      // Redirect after successful update
      setTimeout(() => navigate("/categories"), 2000);
      return { error: null };
    } catch (err) {
      toast.error(err.message || "Failed to update category");
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
          <h1 className="text-2xl font-bold">✏️ Edit Category</h1>
          <Button 
            onClick={() => setShowCategorySearch(!showCategorySearch)}
            className="bg-indigo-600 hover:bg-indigo-700"
          >
            {showCategorySearch ? "Hide Search" : "Find Category"}
          </Button>
        </div>
        
        {showCategorySearch && (
          <CategorySearch 
            onSelectCategory={handleSelectCategory} 
            selectedCategoryId={parseInt(id)}
          />
        )}
        
        {error.error && (
          <div className="bg-red-50 border-l-4 border-red-500 p-4 mb-6">
            <p className="text-red-700">{error.error}</p>
          </div>
        )}

        <form ref={formRef} action={handleSubmit} className="mt-6 space-y-6">
          <div>
            <Label htmlFor="name">Category Name:</Label>
            <Input 
              type="text" 
              id="name" 
              name="name" 
              defaultValue={categoryData.name || ""} 
              placeholder="Enter category name"
              onChange={handleNameChange}
            />
            {error.name && <span className="text-red-600 text-sm block mt-1">{error.name}</span>}
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
              defaultValue={categoryData.slug || ""} 
              placeholder="enter-category-slug"
              readOnly={autoGenerateSlug}
              className={autoGenerateSlug ? "bg-gray-100" : ""}
            />
            {error.slug && <span className="text-red-600 text-sm block mt-1">{error.slug}</span>}
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
            {error.parent_id && <span className="text-red-600 text-sm block mt-1">{error.parent_id}</span>}
            
            {/* Show path information for the selected parent */}
            {categoryData.parent_id && (
              <div className="mt-1 text-sm text-gray-500">
                Full path: {parentCategories.find(p => p.value === String(categoryData.parent_id))?.path || ""}
              </div>
            )}
          </div>

          <div>
            <Label htmlFor="description">Description:</Label>
            {/* <Textarea
              id="description"
              name="description"
              defaultValue={categoryData.description || ""}
              className="mt-1"
              placeholder="Describe this category"
              color="green"
              rows={4}
            /> */}

            <RichTextEditor
              id="description"
              name="description"
              value={categoryData.description || ""} 
              onChange={handleDescriptionChange}
              rows={4}
              maxLength={10000}
              placeholder="Write a detailed product description..."
            />

            {error.description && <span className="text-red-600 text-sm block mt-1">{error.description}</span>}
          </div>

          <div>
            <Label htmlFor="image">Image:</Label>
            {categoryData.image && (
              <div className="mb-2">
                <p className="text-sm text-gray-500 mb-1">Current image:</p>
                <div className="w-24 h-24 border rounded overflow-hidden">
                  <img 
                    src={categoryData.image} 
                    alt={categoryData.name} 
                    className="w-full h-full object-cover"
                  />
                </div>
              </div>
            )}
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
            {error.image && <span className="text-red-600 text-sm block mt-1">{error.image}</span>}
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
              value={categoryData.status || "active"} 
              onChange={(event) => {
                setCategoryData(prev => ({
                  ...prev,
                  status: event.target.value
                }));  
              }}
            />
            {error.status && <span className="text-red-600 text-sm block mt-1">{error.status}</span>}
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
                  Updating...
                </span>
              ) : (
                "✅ Update Category"
              )}
            </Button>
          </div>
        </form>
      </div>
    </>
  );
};

export default Edit;