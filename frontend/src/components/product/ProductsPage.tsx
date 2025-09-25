import { useState, useEffect } from "react";
import type { Product } from "../../types.ts";
import ProductCard from "./ProductCard.tsx";
import { config } from "../../config.ts";

const ProductsPage = () => {
  const [products, setProducts] = useState<Product[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    fetch(`${config.baseApi}/products`)
      .then((res) => {
        if (!res.ok) throw new Error("Loading error");
        return res.json();
      })
      .then((data) => {
        setProducts(data.data);
        setLoading(false);
      })
      .catch((error) => {
        setError(error.message);
        setLoading(false);
      });
  }, []);

  if (loading) return <p>Loading ...</p>;
  if (error) return <p>Error: {error}</p>;

  return (
    <div className="flex flex-wrap justify-center p-4">
      {products.length > 0 ? (
        products.map((product) => (
          <ProductCard key={product.id} product={product} />
        ))
      ) : (
        <p>Products not found</p>
      )}
    </div>
  );
};

export default ProductsPage;
