import { AuthProvider } from "react-admin";

const apiUrl = "http://localhost:3000/api/auth";

const authProvider: AuthProvider = {
  login: async ({ username, password }) => {
    const request = new Request(`${apiUrl}/login`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        name: username,
        password: password,
      })});
    
    let response;
    response = await fetch(request);

    if (response.status < 200 || response.status >= 300) {
      throw new Error(response.statusText);
    }
    let auth = await response.json();
    if (!auth.admin) {
      throw new Error("You are not admin");
    }
    // server set cookies and login is successfull
  },

  checkAuth: async () => {
    const res = await fetch(`${apiUrl}/me`, {
      method: "GET",
      credentials: "include",
    });

    if (res.ok) return Promise.resolve();
    return Promise.reject();
  },

  checkError: async (error) => {
    if (error.status === 401 || error.status === 403) {
      return Promise.reject();
    }
    return Promise.resolve();
  },

  logout: async () => {
    await fetch(`${apiUrl}/logout`, {
      method: "POST",
      credentials: "include",
    });

    return Promise.resolve();
  },

  getIdentity: async () => {
    const res = await fetch(`${apiUrl}/me`, {
      method: "GET",
      credentials: "include",
    });

    if (!res.ok) return Promise.reject();
    const credentials = await res.json();

    const { id, name } = crendentials;

    return { id, name }
  }
}

export default authProvider;
