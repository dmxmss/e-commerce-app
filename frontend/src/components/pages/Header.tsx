import { Link } from "react-router-dom";
import { useCart } from "../cart/CartContext.tsx";
import { useAuth } from "../auth/AuthContext.tsx";

const Header = () => {
  const { cart } = useCart();
  const { user } = useAuth();

  const totalItems = cart.reduce((sum, item) => sum + item.quantity, 0);

  return (
    <header className="bg-white shadow-md p-4 flex justify-between items-center">
      <Link to="/" className="text-2xl font-bold text-green-600">
        MyShop
      </Link>
      <nav className="flex flex-row items-center space-x-6">
        <Link to="/" className="h-fit hover:text-green-600">Main</Link>

        <Link to="/about" className="h-fit hover:text-green-600">About us</Link>

        <Link to="/cart" className="h-fit hover:text-green-600">
          Cart: {totalItems}
        </Link>

        { !user ?
          <>
            <Link
              to="/login"
              className="h-fit px-4 py-1 border text-green-600 rounded-lg hover:bg-green-50 transition"
            >
              Login
            </Link>

            <Link
              to="/signup"
              className="h-fit px-4 py-1 bg-green-600 text-white rounded-lg hover:bg-green-700 transition"
            >
              Sign Up
            </Link> 
          </>
          : 
            <Link to="#" className="flex cursor-pointer items-center space-x-2 bg-green-100 text-green-800 px-3 py-1 rounded-full shadow-sm">
              <div className="w-10 h-10 bg-green-600 text-white rounded-full flex items-center justify-center font-semibold">
                {user.name}
              </div>
              <span className="font-medium">{user.name}</span>
            </Link>
        }
      </nav>
    </header>
  );
};

export default Header;
