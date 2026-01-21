import { createSlice, createAsyncThunk } from "@reduxjs/toolkit";
import { authService } from '../services/authService';


 
export const loginUser = createAsyncThunk("auth/loginUser", async (credentials, thunkAPI) => {
  try {
    const response = await authService.login(credentials.email, credentials.password);
   
    // The response now contains user data without the token (since it's in HttpOnly cookie)
    return response.data.user || response.data;  
  } catch (error) {
    return thunkAPI.rejectWithValue(error.message);
  }
});

export const registerUser = createAsyncThunk("auth/registerUser", async (userData, thunkAPI) => {
  try {
    const response = await authService.register(userData);
   
    return response.data;  
  } catch (error) {
    return thunkAPI.rejectWithValue(error.message);
  }
});

const authSlice = createSlice({
  name: "auth",
  initialState: {
    user: null,
    token: null, // Token is now stored in HttpOnly cookie, not localStorage
    status: "idle",
    error: null,
  },
  reducers: {
    logout: (state) => {
      state.user = null;
      state.token = null;
      // Token is stored in HttpOnly cookie, no need to clear from localStorage here
    },
  },
  extraReducers: (builder) => {
    builder
      .addCase(loginUser.pending, (state) => {
        state.status = "loading";
      })
      .addCase(loginUser.fulfilled, (state, action) => {
        state.status = "succeeded";
        state.user = action.payload.user || action.payload; // Update user data
        // Token is now stored in HttpOnly cookie, no need to store in Redux
      })
      .addCase(loginUser.rejected, (state, action) => {
        state.status = "failed";
        state.error = action.payload;
      })
      .addCase(registerUser.pending, (state) => {
        state.status = "loading";
      })
      .addCase(registerUser.fulfilled, (state, action) => {
        state.status = "succeeded";
        state.user = action.payload; // Update user data
        // Token is now stored in HttpOnly cookie, no need to store in Redux
      })
      .addCase(registerUser.rejected, (state, action) => {
        state.status = "failed";
        state.error = action.payload;
      });
  },
});

export const { logout } = authSlice.actions;
export default authSlice.reducer;
