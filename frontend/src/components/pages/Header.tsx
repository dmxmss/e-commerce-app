import { Link } from "react-router-dom";
import { useCart } from "../cart/CartContext.tsx";

const Header = () => {
  const { cart } = useCart();

  const totalItems = cart.reduce((sum, item) => sum + item.quantity, 0);

  return (
    <header className="bg-white shadow-md p-4 flex justify-between items-center">
      <Link to="/" className="text-2xl font-bold text-green-600">
        MyShop
      </Link>
      <nav className="space-x-6">
        <Link to="/" className="hover:text-green-600">Main</Link>
        <Link to="/about" className="hover:text-green-600">About us</Link>
        <Link to="/cart" className="hover:text-green-600">
          Cart: {totalItems}
        </Link>
      </nav>
    </header>
  );
};

export default Header;
