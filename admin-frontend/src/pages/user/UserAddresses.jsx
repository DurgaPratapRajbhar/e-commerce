import React, { useState, useEffect } from 'react';
import { useAuth } from '../../hooks/useAuth';
import { userService } from '../../services/userService';
import { toast } from 'react-toastify';

const UserAddresses = () => {
  const { user } = useAuth();
  const [addresses, setAddresses] = useState([]);
  const [newAddress, setNewAddress] = useState({
    address_type: 'home',
    street_address: '',
    city: '',
    state: '',
    postal_code: '',
    country: '',
    is_default: false
  });
  const [editingAddress, setEditingAddress] = useState(null);
  const [loading, setLoading] = useState(false);
  const [showForm, setShowForm] = useState(false);

  useEffect(() => {
    if (user?.id) {
      fetchUserAddresses();
    }
  }, [user?.id]);

  const fetchUserAddresses = async () => {
    try {
      setLoading(true);
      const response = await userService.getUserAddresses(user?.id);
      if (response) {
        setAddresses(response);
      }
    } catch (error) {
      console.error('Error fetching addresses:', error);
      toast.error('Failed to fetch addresses');
    } finally {
      setLoading(false);
    }
  };

  const handleInputChange = (e) => {
    const { name, value, type, checked } = e.target;
    setNewAddress(prev => ({
      ...prev,
      [name]: type === 'checkbox' ? checked : value
    }));
  };

  const handleEditChange = (e) => {
    const { name, value, type, checked } = e.target;
    setEditingAddress(prev => ({
      ...prev,
      [name]: type === 'checkbox' ? checked : value
    }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      setLoading(true);
      
      if (editingAddress) {
        // Update existing address
        await userService.updateUserAddress(editingAddress.id, editingAddress);
        toast.success('Address updated successfully');
      } else {
        // Create new address
        const addressData = {
          ...newAddress,
          user_id: user?.id
        };
        await userService.createUserAddress(addressData);
        toast.success('Address created successfully');
      }
      
      setNewAddress({
        address_type: 'home',
        street_address: '',
        city: '',
        state: '',
        postal_code: '',
        country: '',
        is_default: false
      });
      setEditingAddress(null);
      setShowForm(false);
      fetchUserAddresses();
    } catch (error) {
      console.error('Error saving address:', error);
      toast.error('Failed to save address');
    } finally {
      setLoading(false);
    }
  };

  const handleEdit = (address) => {
    setEditingAddress(address);
    setShowForm(true);
  };

  const handleDelete = async (addressId) => {
    if (window.confirm('Are you sure you want to delete this address?')) {
      try {
        setLoading(true);
        await userService.deleteUserAddress(addressId);
        toast.success('Address deleted successfully');
        fetchUserAddresses();
      } catch (error) {
        console.error('Error deleting address:', error);
        toast.error('Failed to delete address');
      } finally {
        setLoading(false);
      }
    }
  };

  const handleSetDefault = async (addressId) => {
    try {
      setLoading(true);
      await userService.setDefaultAddress(user?.id, addressId);
      toast.success('Default address updated successfully');
      fetchUserAddresses();
    } catch (error) {
      console.error('Error setting default address:', error);
      toast.error('Failed to set default address');
    } finally {
      setLoading(false);
    }
  };

  const handleCancel = () => {
    setEditingAddress(null);
    setNewAddress({
      address_type: 'home',
      street_address: '',
      city: '',
      state: '',
      postal_code: '',
      country: '',
      is_default: false
    });
    setShowForm(false);
  };

  if (loading && addresses.length === 0) {
    return (
      <div className="flex justify-center items-center h-64">
        <div className="animate-spin rounded-full h-12 w-12 border-t-2 border-b-2 border-indigo-500"></div>
      </div>
    );
  }

  return (
    <div className="bg-white rounded-xl shadow-lg p-6">
      <div className="flex justify-between items-center mb-6">
        <h2 className="text-2xl font-bold text-gray-800">User Addresses</h2>
        <button
          onClick={() => {
            setEditingAddress(null);
            setShowForm(true);
          }}
          className="px-4 py-2 bg-indigo-600 text-white rounded-lg hover:bg-indigo-700 transition duration-300"
        >
          Add Address
        </button>
      </div>

      {showForm && (
        <form onSubmit={handleSubmit} className="mb-6 p-4 bg-gray-50 rounded-lg">
          <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">
                Address Type
              </label>
              <select
                name="address_type"
                value={editingAddress ? editingAddress.address_type : newAddress.address_type}
                onChange={editingAddress ? handleEditChange : handleInputChange}
                className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500"
              >
                <option value="home">Home</option>
                <option value="work">Work</option>
                <option value="other">Other</option>
              </select>
            </div>
            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">
                Street Address
              </label>
              <input
                type="text"
                name="street_address"
                value={editingAddress ? editingAddress.street_address : newAddress.street_address}
                onChange={editingAddress ? handleEditChange : handleInputChange}
                className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500"
                required
              />
            </div>
            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">
                City
              </label>
              <input
                type="text"
                name="city"
                value={editingAddress ? editingAddress.city : newAddress.city}
                onChange={editingAddress ? handleEditChange : handleInputChange}
                className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500"
                required
              />
            </div>
            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">
                State
              </label>
              <input
                type="text"
                name="state"
                value={editingAddress ? editingAddress.state : newAddress.state}
                onChange={editingAddress ? handleEditChange : handleInputChange}
                className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500"
                required
              />
            </div>
            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">
                Postal Code
              </label>
              <input
                type="text"
                name="postal_code"
                value={editingAddress ? editingAddress.postal_code : newAddress.postal_code}
                onChange={editingAddress ? handleEditChange : handleInputChange}
                className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500"
                required
              />
            </div>
            <div>
              <label className="block text-sm font-medium text-gray-700 mb-1">
                Country
              </label>
              <input
                type="text"
                name="country"
                value={editingAddress ? editingAddress.country : newAddress.country}
                onChange={editingAddress ? handleEditChange : handleInputChange}
                className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-indigo-500 focus:border-indigo-500"
                required
              />
            </div>
            <div className="flex items-center">
              <input
                type="checkbox"
                name="is_default"
                checked={editingAddress ? editingAddress.is_default : newAddress.is_default}
                onChange={editingAddress ? handleEditChange : handleInputChange}
                className="h-4 w-4 text-indigo-600 border-gray-300 rounded"
              />
              <label className="ml-2 block text-sm text-gray-700">
                Set as Default Address
              </label>
            </div>
          </div>

          <div className="flex space-x-3 mt-4">
            <button
              type="submit"
              disabled={loading}
              className="px-4 py-2 bg-indigo-600 text-white rounded-lg hover:bg-indigo-700 disabled:opacity-50 transition duration-300"
            >
              {loading ? 'Saving...' : editingAddress ? 'Update Address' : 'Create Address'}
            </button>
            <button
              type="button"
              onClick={handleCancel}
              className="px-4 py-2 bg-gray-500 text-white rounded-lg hover:bg-gray-600 transition duration-300"
            >
              Cancel
            </button>
          </div>
        </form>
      )}

      <div className="space-y-4">
        {addresses.length === 0 ? (
          <p className="text-gray-500 text-center py-8">No addresses found. Add your first address.</p>
        ) : (
          addresses.map((address) => (
            <div
              key={address.id}
              className={`p-4 border rounded-lg ${
                address.is_default ? 'border-indigo-500 bg-indigo-50' : 'border-gray-200'
              }`}
            >
              <div className="flex justify-between items-start">
                <div className="flex-1">
                  <div className="flex items-center space-x-2 mb-2">
                    <h3 className="font-semibold text-gray-800 capitalize">
                      {address.address_type} Address
                    </h3>
                    {address.is_default && (
                      <span className="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium bg-indigo-100 text-indigo-800">
                        Default
                      </span>
                    )}
                  </div>
                  <p className="text-gray-600">{address.street_address}</p>
                  <p className="text-gray-600">
                    {address.city}, {address.state} {address.postal_code}
                  </p>
                  <p className="text-gray-600">{address.country}</p>
                </div>
                <div className="flex space-x-2">
                  {!address.is_default && (
                    <button
                      onClick={() => handleSetDefault(address.id)}
                      className="text-sm px-3 py-1 bg-indigo-100 text-indigo-700 rounded hover:bg-indigo-200 transition duration-300"
                    >
                      Set as Default
                    </button>
                  )}
                  <button
                    onClick={() => handleEdit(address)}
                    className="text-sm px-3 py-1 bg-blue-100 text-blue-700 rounded hover:bg-blue-200 transition duration-300"
                  >
                    Edit
                  </button>
                  <button
                    onClick={() => handleDelete(address.id)}
                    className="text-sm px-3 py-1 bg-red-100 text-red-700 rounded hover:bg-red-200 transition duration-300"
                  >
                    Delete
                  </button>
                </div>
              </div>
            </div>
          ))
        )}
      </div>
    </div>
  );
};

export default UserAddresses;