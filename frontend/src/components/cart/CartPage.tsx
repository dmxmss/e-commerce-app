import { useCart } from "./CartContext.tsx";
import { config } from "../../config.ts";

const CartPage = () => {
  const { cart, addToCart, removeOneFromCart, clearCart, createPayment } = useCart();

  if (cart.length === 0) {
    return (
      <div className="max-w-4xl mx-auto p-6 text-center">
        <h2 className="text-3xl font-bold mb-4">Your cart it empty</h2>
        <p className="text-gray-600">Add some products</p>
      </div> 
    );
  }

  const totalPrice = cart.reduce((sum, item) => sum + item.price * item.quantity, 0);

  return (
    <div className="max-w-4xl mx-auto p-6 bg-white shadow-lg rounded-2xl mt-6">
      <h2 className="text-3xl font-bold mb-6">Your cart</h2>

      <div className="space-y-4">
        {cart.map((item) => ( 
          <div 
            key={item.id}
            className="flex justify-between items-center border-b pb-4 last:border-none"
          >
            <div className="flex items-center gap-4">
              <img src={`${config.imageServer}/${item.images[0]}`} alt={item.name} className="w-20 h-20 bg-gray-300 object-cover rounded-lg" />
              <div>
                <h3 className="text-lg font-semibold">{item.name}</h3>
                <p className="text-gray-500">${item.price.toFixed(2)}</p>
              </div>
            </div>

            <div className="flex items-center gap-4">
              <button
                className="px-3 py-1 w-8 h-8 bg-gray-200 rounded-lg hover:bg-gray-300 transition"
                onClick={() => addToCart({ id: item.id, name: item.name, price: item.price, images: item.images })}
              >
                +
              </button>
              <span className="min-w-[2rem] text-center">{item.quantity}</span>
              <button
                className="px-3 py-1 w-8 h-8 bg-gray-200 rounded-lg hover:bg-gray-300 transition"
                onClick={() => removeOneFromCart(item.id)}
              >
                -
              </button>
            </div>
          </div>  
        ))}
      </div>

      <div className="flex mt-6 justify-between items-center">
        <p className="text-xl font-bold">
          Total: <span className="text-green-600">${totalPrice.toFixed(2)}</span>
        </p>
      
        <div>
          <button 
            className="bg-green-500 text-white px-6 py-2 rounded-xl hover:bg-green-600 transition mr-4"
            onClick={createPayment}
          >
            Buy
          </button>

          <button 
            className="bg-red-500 text-white px-6 py-2 rounded-xl hover:bg-red-600 transition"
            onClick={clearCart}
          >
            Clear cart
          </button>
        </div>
      </div>
    </div>
  );
}

export default CartPage;
