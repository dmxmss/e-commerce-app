import React from "react";
import type { Product } from "./types.ts";
import { Link } from "react-router-dom";
import { useCart } from "../cart/CartContext.tsx";
import { config } from "../../config.ts";

interface ProductCardProps {
  product: Product;
}

const ProductCard: React.FC<ProductCardProps> = ({ product }) => {
  const { addToCart } = useCart();
  return (
    <div className="bg-white rounded-xl shadow-md overflow-hidden hover:shadow-xl transition-shadow duration-300 w-64 m-4 flex flex-col">
      <div className="flex flex-col flex-1">
        <Link key={product.id} to={`/products/${product.id}`}>
          <img src={`${config.imageServer}/${product.images[0]}`} alt={product.name} className="h-48 w-full object-contain bg-gray-100 p-4 mb-2" />
        </Link>
        <div className="p-4">
          <h3 className="text-lg font-semibold text-gray-800 mb-2 line-clamp-2">
            {product.name}
          </h3>
          <p className="text-gray-600 mb-4">${product.price}</p>
          <button className="mt-auto bg-green-500 text-white font-medium py-2 rounded-lg hover:bg-green-600 transition-colors" onClick={() => addToCart(product)}>
            To cart
          </button>
        </div>
      </div>
    </div>
  );
};

export default ProductCard;
