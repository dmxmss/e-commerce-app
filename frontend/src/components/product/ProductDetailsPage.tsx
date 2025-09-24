import { useEffect, useState } from "react";
import { useParams, Link } from "react-router-dom";
import type { Product } from "./types.ts";
import { useCart } from "../cart/CartContext.tsx";

const ProductDetailsPage = () => {
  const { addToCart } = useCart();

  const { id } = useParams<{id: string}>();
  const [product, setProduct] = useState<Product | null>(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    if (!id) return;
    
    fetch(`http://localhost:3001/api/products/${id}`)
      .then((res) => {
        if (!res.ok) throw new Error("Loading error");
        return res.json();
      })
      .then((data) => {
        setProduct(data);
        setLoading(false);
      })
      .catch((error) => {
        setError(error.message);
        setLoading(false);
      });

  }, [id]);

  if (loading) return <p>Loading ...</p>;
  if (error) return <p>Error: {error}</p>;
  if (!product) return <p>Product not found</p>;

  return (
    <div className="min-w-screen min-h-screen bg-gray-50 flex flex-col">
      <header className="p-6 bg-white shadow-md">
        <Link to="/" className="text-blue-600 hover:underline text-lg">
          ← Back
        </Link>
      </header>

      <main className="flex flex-col lg:flex-row flex-1 max-w-6xl mx-auto w-full bg-white shadow-lg rounded-xl overflow-hidden my-8">
        <div className="flex-1 flex items-center justify-center bg-gray-100">
          <img
            src={`http://localhost:3002/${product.images[0]}`}
            alt={product.name}
            className="max-h-[500px] object-contain"
          />
        </div>

        <div className="flex-1 p-8 flex flex-col justify-between">
          <div>
            <h1 className="text-3xl font-bold text-gray-800 mb-4">{product.name}</h1>
            <p className="text-gray-700 mb-6 text-lg leading-relaxed">
              {product.description}
            </p>
            <p className="text-2xl font-semibold text-green-600">${product.price}</p>
          </div>

          <button className="mt-8 bg-green-500 text-white font-medium text-lg py-3 rounded-xl hover:bg-green-600 transition-colors" 
                  onClick={() => addToCart(product)}
          >
            Add to cart
          </button>
        </div>
      </main>
    </div>
  )
}

export default ProductDetailsPage;
