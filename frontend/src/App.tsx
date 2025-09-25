import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import ProductDetailsPage from "./components/product/ProductDetailsPage.tsx";
import Layout from "./Layout.tsx";
import Dashboard from "./components/pages/Dashboard.tsx"
import CartPage from "./components/cart/CartPage.tsx";
import PaymentPage from "./components/pages/PaymentPage.tsx";
import NotFoundPage from "./components/pages/NotFoundPage.tsx";

import CartProvider from "./components/cart/CartContext.tsx"

function App() {
  return (
    <Router>
      <CartProvider>
        <Routes>
          <Route path="/" element={<Layout />}>
            <Route index element={<Dashboard />} />
            <Route path="/products/:id" element={<ProductDetailsPage />} />
            <Route path="/cart" element={<CartPage />} />
            <Route path="/checkout" element={<PaymentPage />} />
            <Route path="*" element={<NotFoundPage />} />
          </Route>
        </Routes>
      </CartProvider>
    </Router>
  )
}

export default App
