import React, { createContext, useContext, useState, useEffect } from "react";

export type CartItem = {
  id: number;
  name: string;
  price: number;
  quantity: number;
};

type CartContextType = {
  cart: CartItem[];
  addToCart: (item: Omit<CartItem, "quantity">) => void;
  removeOneFromCart: (id: number) => void;
  removeFromCart: (id: number) => void;
  clearCart: () => void;
};

const CartContext = createContext<CartContextType | undefined>(undefined);

const CartProvider: React.FC<{ children: React.ReactNode }> = ({children}) => {
  const [cart, setCart] = useState<CartItem[]>(() => {
    const savedCart = localStorage.getItem("cart");
    return savedCart ? JSON.parse(savedCart) : [];
  });

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

  return (
    <CartContext.Provider value={{ cart, addToCart, removeFromCart, removeOneFromCart, clearCart }}>
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
