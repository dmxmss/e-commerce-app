import React, { createContext, useContext, useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { config } from "../../config.ts";
import type { CartItem } from "../../types.ts";

type CartContextType = {
  cart: CartItem[];
  addToCart: (item: Omit<CartItem, "quantity">) => void;
  removeOneFromCart: (id: number) => void;
  removeFromCart: (id: number) => void;
  clearCart: () => void;
  createPayment: () => void;
};

const CartContext = createContext<CartContextType | undefined>(undefined);

const CartProvider: React.FC<{ children: React.ReactNode }> = ({children}) => {
  const [cart, setCart] = useState<CartItem[]>(() => {
    const savedCart = localStorage.getItem("cart");
    return savedCart ? JSON.parse(savedCart) : [];
  });

  const navigate = useNavigate();

  useEffect(() => {
    localStorage.setItem("cart", JSON.stringify(cart));
  }, [cart]);

  const addToCart = (item: Omit<CartItem, "quantity">) => {
    setCart(prev => 
      prev.some(cartItem => cartItem.id === item.id) // increment item quantity if it is present in the cart and add new item if it is not
        ? prev.map(cartItem => 
          cartItem.id === item.id
           ? { ...cartItem, quantity: cartItem.quantity + 1}
           : cartItem
          )
        : [ ...prev, { ...item, quantity: 1 }]
    );
  };

  const removeOneFromCart = (id: number) => {
    setCart(prev => 
      prev.some(cartItem => cartItem.id === id)
        ? prev.map(cartItem =>
          cartItem.id === id
            ? cartItem.quantity > 0
              ? { ...cartItem, quantity: cartItem.quantity - 1 }
              : cartItem 
            : cartItem 
          )
        : prev
    );
  };

  const removeFromCart = (id: number) => {
    setCart(prev => prev.filter(cartItem => cartItem.id !== id));
  };

  const clearCart = () => setCart([]);

  const createPayment = () => {
    const products = cart.filter((item) => item.quantity > 0);
    if (products.length == 0) {
      return;
    }

    fetch(`${config.baseApi}/payments`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        product_ids: products.map((item) => item.id),
        currency: "USD",
      })
    })
      .then((res) => {
        if (!res.ok) throw new Error(res.statusText);
        return res.json();
    })
      .then((data) => {
        navigate("checkout", { state: { clientSecret: data.client_secret }});
    })
  };

  return (
    <CartContext.Provider value={{ cart, addToCart, removeFromCart, removeOneFromCart, clearCart, createPayment }}>
      {children}
    </CartContext.Provider>
  )
};

export const useCart = () => {
  const context = useContext(CartContext);
  if (!context) {
    throw new Error("useCart must be used inside CartProvider");
  }
  return context;
};

export default CartProvider;
