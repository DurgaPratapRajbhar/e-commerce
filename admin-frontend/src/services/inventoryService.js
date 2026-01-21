import { api } from '../lib/api';
import { API_ENDPOINTS } from '../config/apiConfig';

export const inventoryService = {
  // Inventory Management
  getInventory: () => api.get(API_ENDPOINTS.INVENTORY.BASE),
  getLowStockItems: () => api.get(API_ENDPOINTS.INVENTORY.LOW_STOCK),
  getInventoryByProductVariant: (productId, variantId) => api.get(API_ENDPOINTS.INVENTORY.BY_PRODUCT_VARIANT(productId, variantId)),
  createInventory: (inventoryData) => api.post(API_ENDPOINTS.INVENTORY.BASE, inventoryData),
  updateInventory: (inventoryId, inventoryData) => api.put(API_ENDPOINTS.INVENTORY.BASE, inventoryData),
  deleteInventory: (inventoryId) => api.delete(API_ENDPOINTS.INVENTORY.BASE),

  // Inventory Transactions
  getTransactions: () => api.get(API_ENDPOINTS.INVENTORY.TRANSACTIONS),
  getTransactionsByProduct: (productId) => api.get(API_ENDPOINTS.INVENTORY.TRANSACTIONS_BY_PRODUCT(productId)),
  getTransactionsByVariant: (productId, variantId) => api.get(API_ENDPOINTS.INVENTORY.TRANSACTIONS_BY_VARIANT(productId, variantId)),
  getRecentTransactions: () => api.get(API_ENDPOINTS.INVENTORY.TRANSACTIONS_RECENT),
  getTransactionsByReference: (refId) => api.get(API_ENDPOINTS.INVENTORY.TRANSACTIONS_BY_REFERENCE(refId)),
  createTransaction: (transactionData) => api.post(API_ENDPOINTS.INVENTORY.TRANSACTIONS, transactionData),
};