import { api } from '../lib/api';
import { API_ENDPOINTS } from '../config/apiConfig';

export const shipmentService = {
  // Shipment Management
  getShipments: () => api.get(API_ENDPOINTS.SHIPMENT.ALL),
  getShipmentById: (shipmentId) => api.get(API_ENDPOINTS.SHIPMENT.BY_ID(shipmentId)),
  getShipmentsByOrder: (orderId) => api.get(API_ENDPOINTS.SHIPMENT.BY_ORDER(orderId)),
  createShipment: (shipmentData) => api.post(API_ENDPOINTS.SHIPMENT.BASE, shipmentData),
  updateShipment: (shipmentId, shipmentData) => api.put(API_ENDPOINTS.SHIPMENT.BY_ID(shipmentId), shipmentData),
  updateShipmentStatus: (shipmentId, statusData) => api.put(API_ENDPOINTS.SHIPMENT.UPDATE_STATUS(shipmentId), statusData),
  getShipmentStatus: (shipmentId) => api.get(API_ENDPOINTS.SHIPMENT.STATUS(shipmentId)),
  deleteShipment: (shipmentId) => api.delete(API_ENDPOINTS.SHIPMENT.BY_ID(shipmentId)),

  // Tracking Management
  getTrackingEvents: (shipmentId) => api.get(API_ENDPOINTS.SHIPMENT.TRACKING_BY_ID(shipmentId)),
  getLatestTrackingEvent: (shipmentId) => api.get(API_ENDPOINTS.SHIPMENT.TRACKING_LATEST(shipmentId)),
  createTrackingEvent: (trackingData) => api.post(API_ENDPOINTS.SHIPMENT.TRACKING, trackingData),
};