import React, { createContext, useContext, useState, useEffect, useCallback } from "react";
import type { User } from "../../types.ts";
import { config } from "../../config.ts";

type AuthContextType = {
  user: User | null;
  checkAuth: () => void;
  logout: () => void;
};

const AuthContext = createContext<AuthContextType | undefined>(undefined);

const AuthProvider: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  const [ user, setUser ] = useState<User | null>(null);

  const logout = useCallback(async () => {
    await fetch(`${config.baseApi}/auth/logout`, {
      method: "POST",
    });
    setUser(null);
  }, [])

  const checkAuth = useCallback(async () => {
    await fetch(`${config.baseApi}/auth/me`)
      .then(async (res) => {
        if (res.ok) {
          const data = await res.json() as User;
          setUser(data);
        } else if (res.status === 401) {
          logout();
        }
      })
  }, [logout])

  useEffect(() => { // checking auth with interval
    checkAuth();

    const interval = setInterval(checkAuth, 5*60*1000);
    return () => clearInterval(interval);
  }, [checkAuth]);

  return (
    <AuthContext.Provider value={{ user, checkAuth, logout }}>
      {children}
    </AuthContext.Provider>
  )
};

export const useAuth = () => {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error("useAuth must be used inside AuthProvider");
  }

  return context;
}

export default AuthProvider;
