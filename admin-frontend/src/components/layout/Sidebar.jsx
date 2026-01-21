import React, { useState } from 'react';
import { Link } from "react-router-dom";

const Sidebar = () => {
  const [accordions, setAccordions] = useState({
    userManagement: false,
    productManagement: false,
    orderManagement: false
  });

  const toggleAccordion = (key) => {
    setAccordions(prev => ({
      ...prev,
      [key]: !prev[key]
    }));
  };

  return (
    <div className="w-64 bg-slate-900 text-white fixed left-0 top-0 h-full shadow-xl overflow-y-auto">
      {/* Header */}
      <div className="px-6 py-4 bg-slate-800 border-b border-slate-700">
        <h2 className="text-base font-semibold text-white">E-Commerce Admin</h2>
      </div>

      {/* Navigation */}
      <nav className="px-3 py-4 space-y-1">
        <NavItem to="/dashboard" label="Dashboard" icon="home" />
        
        {/* User Management Accordion */}
        <AccordionItem 
          label="User Management" 
          icon="user"
          isOpen={accordions.userManagement}
          onToggle={() => toggleAccordion('userManagement')}
        >
          <SubNavItem to="/user/profile" label="User Profile" />
          <SubNavItem to="/user/addresses" label="User Addresses" />
          <SubNavItem to="/user/list" label="User List" />
          <SubNavItem to="/user/addresses/list" label="Address List" />
        </AccordionItem>

        {/* Product Management Accordion */}
        <AccordionItem 
          label="Product Management" 
          icon="box"
          isOpen={accordions.productManagement}
          onToggle={() => toggleAccordion('productManagement')}
        >
          <SubNavItem to="/categories" label="Categories" />
          <SubNavItem to="/products" label="Products" />
          <SubNavItem to="/product-images" label="Product Images" />
          <SubNavItem to="/product-reviews" label="Reviews" />
          <SubNavItem to="/product-units" label="Units (UOM)" />
        </AccordionItem>

        {/* Order Management Accordion */}
        <AccordionItem 
          label="Order Management" 
          icon="cart"
          isOpen={accordions.orderManagement}
          onToggle={() => toggleAccordion('orderManagement')}
        >
          <SubNavItem to="/carts" label="Shopping Carts" />
          <SubNavItem to="/orders" label="Orders" />
          <SubNavItem to="/shipments" label="Shipments" />
          <SubNavItem to="/payments" label="Payments" />
          <SubNavItem to="/refunds" label="Refunds" />
        </AccordionItem>

        <NavItem to="/inventory" label="Inventory" icon="package" />
      </nav>
    </div>
  );
};

// NavItem Component
const NavItem = ({ to, label, icon }) => {
  const icons = {
    home: "M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6",
    user: "M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z",
    box: "M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4",
    cart: "M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z",
    package: "M20 13V6a2 2 0 00-2-2H6a2 2 0 00-2 2v7m16 0v5a2 2 0 01-2 2H6a2 2 0 01-2-2v-5m16 0h-2.586a1 1 0 00-.707.293l-2.414 2.414a1 1 0 01-.707.293h-3.172a1 1 0 01-.707-.293l-2.414-2.414A1 1 0 006.586 13H4",
    settings: "M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z M15 12a3 3 0 11-6 0 3 3 0 016 0z",
  };

  return (
    <Link 
      to={to} 
      className="flex items-center px-3 py-2.5 text-sm text-slate-300 hover:bg-slate-800 hover:text-white rounded-md transition-colors duration-150"
    >
      <svg className="w-4 h-4 mr-3 flex-shrink-0" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
        <path d={icons[icon]}></path>
      </svg>
      <span className="text-sm font-medium truncate">{label}</span>
    </Link>
  );
};

// Accordion Item Component
const AccordionItem = ({ label, icon, isOpen, onToggle, children }) => {
  const icons = {
    user: "M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z",
    box: "M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4",
    cart: "M3 3h2l.4 2M7 13h10l4-8H5.4M7 13L5.4 5M7 13l-2.293 2.293c-.63.63-.184 1.707.707 1.707H17m0 0a2 2 0 100 4 2 2 0 000-4zm-8 2a2 2 0 11-4 0 2 2 0 014 0z",
  };

  return (
    <div className="mb-1">
      <button
        onClick={onToggle}
        className={`flex items-center justify-between w-full px-3 py-2.5 text-sm rounded-md transition-colors duration-150 ${
          isOpen 
            ? 'bg-slate-800 text-white' 
            : 'text-slate-300 hover:bg-slate-800 hover:text-white'
        }`}
      >
        <div className="flex items-center min-w-0">
          <svg className="w-4 h-4 mr-3 flex-shrink-0" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
            <path d={icons[icon]}></path>
          </svg>
          <span className="text-sm font-medium truncate">{label}</span>
        </div>
        <svg 
          className={`w-4 h-4 ml-2 flex-shrink-0 transition-transform duration-200 ${isOpen ? 'rotate-180' : ''}`} 
          viewBox="0 0 24 24" 
          fill="none" 
          stroke="currentColor" 
          strokeWidth="2"
          strokeLinecap="round" 
          strokeLinejoin="round"
        >
          <path d="M19 9l-7 7-7-7"></path>
        </svg>
      </button>
      
      {isOpen && (
        <div className="mt-1 ml-7 pl-3 border-l border-slate-700 space-y-0.5">
          {children}
        </div>
      )}
    </div>
  );
};

// Sub Navigation Item Component
const SubNavItem = ({ to, label }) => {
  return (
    <Link 
      to={to} 
      className="block py-2 px-3 text-xs text-slate-400 hover:text-white hover:bg-slate-800 rounded-md transition-colors duration-150"
    >
      {label}
    </Link>
  );
};

export default Sidebar;