const Dashboard = () => {
  return (
    <div className="flex-1 p-6 bg-gradient-to-br from-green-100 to-green-300 min-h-screen">
      {/* Welcome Message */}
      <h2 className="text-3xl font-bold text-gray-800 mb-6">📊 Dashboard Overview</h2>
      <p className="text-lg text-gray-600 mb-8">Welcome back! Here’s what’s happening in your application today.</p>

      {/* Stats Cards */}
      <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
        <div className="bg-white shadow-lg rounded-xl p-6 flex flex-col items-center">
          <div className="text-blue-500 text-5xl font-bold">1,245</div>
          <p className="text-gray-600 mt-2">Total Users</p>
        </div>

        <div className="bg-white shadow-lg rounded-xl p-6 flex flex-col items-center">
          <div className="text-green-500 text-5xl font-bold">3,578</div>
          <p className="text-gray-600 mt-2">Total Orders</p>
        </div>

        <div className="bg-white shadow-lg rounded-xl p-6 flex flex-col items-center">
          <div className="text-yellow-500 text-5xl font-bold">$45,720</div>
          <p className="text-gray-600 mt-2">Revenue</p>
        </div>

        <div className="bg-white shadow-lg rounded-xl p-6 flex flex-col items-center">
          <div className="text-red-500 text-5xl font-bold">438</div>
          <p className="text-gray-600 mt-2">New Signups</p>
        </div>
      </div>

      {/* Activity Table */}
      <div className="mt-10 bg-white shadow-lg rounded-xl p-6">
        <h3 className="text-xl font-semibold text-gray-800 mb-4">📝 Recent Activity</h3>
        <table className="w-full border-collapse border border-gray-300">
          <thead className="bg-gray-200">
            <tr>
              <th className="border border-gray-300 px-4 py-2">User</th>
              <th className="border border-gray-300 px-4 py-2">Action</th>
              <th className="border border-gray-300 px-4 py-2">Date</th>
            </tr>
          </thead>
          <tbody>
            <tr className="text-center">
              <td className="border border-gray-300 px-4 py-2">John Doe</td>
              <td className="border border-gray-300 px-4 py-2">Purchased a plan</td>
              <td className="border border-gray-300 px-4 py-2">Feb 8, 2025</td>
            </tr>
            <tr className="text-center bg-gray-100">
              <td className="border border-gray-300 px-4 py-2">Jane Smith</td>
              <td className="border border-gray-300 px-4 py-2">Signed up</td>
              <td className="border border-gray-300 px-4 py-2">Feb 7, 2025</td>
            </tr>
            <tr className="text-center">
              <td className="border border-gray-300 px-4 py-2">Michael Brown</td>
              <td className="border border-gray-300 px-4 py-2">Updated profile</td>
              <td className="border border-gray-300 px-4 py-2">Feb 6, 2025</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  );
};

export default Dashboard;
