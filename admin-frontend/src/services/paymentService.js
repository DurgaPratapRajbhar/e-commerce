import { api } from '../lib/api';
import { API_ENDPOINTS } from '../config/apiConfig';

export const paymentService = {
  // Payment Management
  getPayments: () => api.get(API_ENDPOINTS.PAYMENT.BASE),
  getPaymentById: (paymentId) => api.get(API_ENDPOINTS.PAYMENT.BY_ID(paymentId)),
  getPaymentsByOrder: (orderId) => api.get(API_ENDPOINTS.PAYMENT.BY_ORDER(orderId)),
  getPaymentsByStatus: (status) => api.get(API_ENDPOINTS.PAYMENT.BY_STATUS(status)),
  createPayment: (paymentData) => api.post(API_ENDPOINTS.PAYMENT.BASE, paymentData),
  updatePayment: (paymentId, paymentData) => api.put(API_ENDPOINTS.PAYMENT.BY_ID(paymentId), paymentData),
  updatePaymentStatus: (paymentId, statusData) => api.put(API_ENDPOINTS.PAYMENT.UPDATE_STATUS(paymentId), statusData),
  deletePayment: (paymentId) => api.delete(API_ENDPOINTS.PAYMENT.BY_ID(paymentId)),

  // Refund Management
  getRefunds: () => api.get(API_ENDPOINTS.PAYMENT.REFUNDS),
  getRefundById: (refundId) => api.get(API_ENDPOINTS.PAYMENT.REFUND_BY_ID(refundId)),
  getRefundsByOrder: (orderId) => api.get(API_ENDPOINTS.PAYMENT.REFUNDS_BY_ORDER(orderId)),
  getRefundsByStatus: (status) => api.get(API_ENDPOINTS.PAYMENT.REFUNDS_BY_STATUS(status)),
  createRefund: (refundData) => api.post(API_ENDPOINTS.PAYMENT.REFUNDS, refundData),
  updateRefund: (refundId, refundData) => api.put(API_ENDPOINTS.PAYMENT.REFUND_BY_ID(refundId), refundData),
  deleteRefund: (refundId) => api.delete(API_ENDPOINTS.PAYMENT.REFUND_BY_ID(refundId)),
};