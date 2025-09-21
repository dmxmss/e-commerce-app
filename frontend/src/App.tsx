import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import ProductDetailsPage from "./components/product/ProductDetailsPage.tsx";
import Layout from "./Layout.tsx";
import Dashboard from "./components/pages/Dashboard.tsx"
import CartPage from "./components/cart/CartPage.tsx";
import NotFoundPage from "./components/pages/NotFoundPage.tsx";

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Layout />}>
          <Route index element={<Dashboard />} />
          <Route path="/products/:id" element={<ProductDetailsPage />} />
          <Route path="/cart" element={<CartPage />} />
          <Route path="*" element={<NotFoundPage />} />
        </Route>
      </Routes>
    </Router>
  )
}

export default App
