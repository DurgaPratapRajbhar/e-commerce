import { useState, useEffect, useRef } from "react";
import { useNavigate } from "react-router-dom";
import { PencilSquareIcon, TrashIcon, EyeIcon } from "@heroicons/react/24/solid";
import { api } from "../../lib/api";
import Pagination from "../../components/Pagination";
import ProductFilter from "./ProductFilter";
import ProductDetailModal from "./ProductDetailModal";

const ProductList = () => {
  const env = import.meta.env;
  const productUrl = env.VITE_PRODUCT_SERVICE;
  const productEndpoint = env.VITE_PRODUCT_LIST;

  const [products, setProducts] = useState([]);
  const [currentPage, setCurrentPage] = useState(1);
  const [totalPages, setTotalPages] = useState(1);
  const [totalItems, setTotalItems] = useState(0);
  const [itemsPerPage] = useState(10);
  const [isLoading, setIsLoading] = useState(true);
  const [filters, setFilters] = useState({
    search: "",
    category: "",
    status: "",
    priceMin: "",
    priceMax: ""
  });
  const [sortField, setSortField] = useState("id");
  const [sortDirection, setSortDirection] = useState("asc");
  const [selectedProduct, setSelectedProduct] = useState(null);
  const [showDetailModal, setShowDetailModal] = useState(false);
  const handledUseRef = useRef(false);
  const navigate = useNavigate();

  useEffect(() => {
    if (!handledUseRef.current) {
      fetchProducts();
      handledUseRef.current = true;
    }
  }, []);

  useEffect(() => {
    // When filters or sort options change, fetch products
    if (handledUseRef.current) {
      fetchProducts(1);
    }
  }, [filters, sortField, sortDirection]);

  const buildQueryParams = (page) => {
    let params = new URLSearchParams();
    params.append("page", page);
    params.append("limit", itemsPerPage);
    params.append("sort", sortField);
    params.append("direction", sortDirection);
    
    if (filters.search) params.append("search", filters.search);
    if (filters.category) params.append("category", filters.category);
    if (filters.status) params.append("status", filters.status);
    if (filters.priceMin) params.append("price_min", filters.priceMin);
    if (filters.priceMax) params.append("price_max", filters.priceMax);
    
    return params.toString();
  };

  const fetchProducts = async (page = 1) => {
    setIsLoading(true);
    try {
      const queryParams = buildQueryParams(page);
      const response = await api.get(`${productUrl}${productEndpoint}?${queryParams}`);
      const { data, total } = response.data;
      setProducts(data);
      setTotalItems(total);
      setTotalPages(Math.ceil(total / itemsPerPage));
      setCurrentPage(page);
    } catch (err) {
      console.error("Error fetching products:", err);
    } finally {
      setIsLoading(false);
    }
  };

  const handlePageChange = (page) => {
    if (page >= 1 && page <= totalPages) {
      fetchProducts(page);
    }
  };

  const handleFilterChange = (newFilters) => {
    setFilters(newFilters);
  };

  const handleSort = (field) => {
    if (field === sortField) {
      setSortDirection(sortDirection === "asc" ? "desc" : "asc");
    } else {
      setSortField(field);
      setSortDirection("asc");
    }
  };

  const viewProductDetails = (product) => {
    setSelectedProduct(product);
    setShowDetailModal(true);
  };

  const closeDetailModal = () => {
    setShowDetailModal(false);
    setSelectedProduct(null);
  };

  const deleteProduct = async (id) => {
    if (window.confirm("Are you sure you want to delete this product?")) {
      try {
        await api.delete(`${productUrl+productEndpoint}/${id}`);
        fetchProducts(currentPage);
      } catch (err) {
        console.error("Error deleting product:", err);
      }
    }
  };

  const formatPrice = (price, discount = 0) => {
    const actualPrice = price - (price * discount / 100);
    return new Intl.NumberFormat('en-US', { style: 'currency', currency: 'INR' }).format(actualPrice);
  };

  // Sort indicator function
  const getSortIndicator = (field) => {
    if (field !== sortField) return null;
    return sortDirection === "asc" ? "↑" : "↓";
  };

  return (
    <div className="container mx-auto px-4 py-8">
      <div className="bg-white shadow-xl rounded-xl p-6">
        <div className="flex justify-between items-center mb-6">
          <h2 className="text-2xl font-bold text-gray-800 flex items-center gap-2">
            <span>📦</span> Products
            {!isLoading && <span className="text-sm font-normal text-gray-500 ml-2">({totalItems} items)</span>}
          </h2>
          <button
            className="bg-indigo-600 text-white px-4 py-2 rounded-lg hover:bg-indigo-700 transition duration-200 flex items-center gap-2"
            onClick={() => navigate("/products/add")}
          >
            <span>➕</span> Add Product
          </button>
        </div>

        {/* Filter Component */}
        <ProductFilter filters={filters} onFilterChange={handleFilterChange} />

        <div className="overflow-x-auto mt-6">
          {isLoading ? (
            <div className="flex justify-center items-center h-64">
              <div className="animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-indigo-500"></div>
            </div>
          ) : products.length === 0 ? (
            <div className="text-center py-12 text-gray-500">
              No products found. Try adjusting your filters.
            </div>
          ) : (
            <table className="w-full border-collapse">
              <thead className="bg-gray-100 text-gray-600">
                <tr>
                  {[
                    { key: "#", sortable: false },
                    { key: "Image", sortable: false },
                    { key: "Name", field: "name", sortable: true },
                    { key: "Price", field: "price", sortable: true },
                    { key: "Discount", field: "discount", sortable: true },
                    { key: "Final Price", sortable: false },
                    { key: "Brand", field: "brand", sortable: true },
                    { key: "Category", field: "category_id", sortable: true },
                    { key: "Stock", field: "stock", sortable: true },
                    { key: "Status", field: "status", sortable: true },
                    { key: "Actions", sortable: false }
                  ].map((header) => (
                    <th 
                      key={header.key} 
                      className={`px-4 py-3 text-left text-sm font-semibold border-b ${header.sortable ? 'cursor-pointer hover:bg-gray-200' : ''}`}
                      onClick={() => header.sortable && handleSort(header.field)}
                    >
                      {header.key} {header.sortable && getSortIndicator(header.field)}
                    </th>
                  ))}
                </tr>
              </thead>
              <tbody className="divide-y divide-gray-200">
                {products.map((product, index) => (
                  <tr key={product.id} className="hover:bg-gray-50 transition duration-150">
                    <td className="px-4 py-3 text-sm">{(currentPage - 1) * itemsPerPage + index + 1}</td>
                    <td className="px-4 py-3">

                      {/* {product.primary_image ? (
                        <img 
                          src={product.primary_image} 
                          alt={product.name} 
                          className="h-10 w-10 object-cover rounded"
                        />
                      ) : (
                        <div className="h-10 w-10 bg-gray-200 rounded flex items-center justify-center">
                          <span className="text-xs text-gray-500">No image</span>
                        </div>
                      )} */}

                      <button
                        className="h-10 w-10 flex items-center justify-center bg-gray-200 hover:bg-gray-300 rounded-lg transition"
                        onClick={() => navigate(`/products/${product.id}/product-images`)}
                      >
                        🖼️
                      </button>


                    </td>

                    <td className="px-4 py-3 text-sm font-medium">{product.name}</td>
                    <td className="px-4 py-3 text-sm">{formatPrice(product.price, 0)}</td>
                    <td className="px-4 py-3 text-sm">
                      {product.discount > 0 ? (
                        <span className="px-2 py-1 bg-red-100 text-red-800 rounded-full text-xs">
                          {product.discount}%
                        </span>
                      ) : (
                        <span className="text-gray-400">-</span>
                      )}
                    </td>
                    <td className="px-4 py-3 text-sm font-medium">
                      {formatPrice(product.price, product.discount)}
                      {product.discount > 0 && (
                        <span className="text-xs text-gray-500 line-through ml-2">
                          {formatPrice(product.price)}
                        </span>
                      )}
                    </td>
                    <td className="px-4 py-3 text-sm">{product.brand || "N/A"}</td>
                    <td className="px-4 py-3 text-sm">{product.category?.name || "N/A"}</td>
                    <td className="px-4 py-3 text-sm">
                      <span className={`${product.stock > 10 ? 'text-green-600' : product.stock > 0 ? 'text-yellow-600' : 'text-red-600'}`}>
                        {product.stock}
                        {product.unit && product.quantity_value ? ` ${product.unit.symbol}` : ''}
                      </span>
                    </td>
                    <td className="px-4 py-3 text-sm">
                      <span className={`px-2 py-1 rounded-full text-xs ${
                        product.status === 'active' 
                          ? 'bg-green-100 text-green-800' 
                          : product.status === 'draft' 
                            ? 'bg-yellow-100 text-yellow-800' 
                            : 'bg-red-100 text-red-800'
                      }`}>
                        {product.status}
                      </span>
                    </td>
                    <td className="px-4 py-3 text-sm">
                      <div className="flex items-center gap-3">
                        <button 
                          className="text-blue-600 hover:text-blue-800" 
                          onClick={() => viewProductDetails(product)}
                          title="View details"
                        >
                          <EyeIcon className="h-5 w-5" />
                        </button>
                        <button 
                          className="text-indigo-600 hover:text-indigo-800" 
                          onClick={() => navigate(`/products/edit/${product.id}`)}
                          title="Edit product"
                        >
                          <PencilSquareIcon className="h-5 w-5" />
                        </button>
                        <button 
                          className="text-red-600 hover:text-red-800" 
                          onClick={() => deleteProduct(product.id)}
                          title="Delete product"
                        >
                          <TrashIcon className="h-5 w-5" />
                        </button>
                      </div>
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          )}
        </div>

        {/* Pagination Component */}
        {totalPages > 1 && (
          <div className="mt-6">
            <Pagination
              currentPage={currentPage}
              totalPages={totalPages}
              onPageChange={handlePageChange}
            />
          </div>
        )}
      </div>

      {/* Product Detail Modal */}
      {showDetailModal && selectedProduct && (
        <ProductDetailModal 
          product={selectedProduct} 
          onClose={closeDetailModal} 
        />
      )}
    </div>
  );
};

export default ProductList;