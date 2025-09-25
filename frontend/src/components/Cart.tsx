import React from "react";

interface Food {
  _id: string;
  name: string;
  price: number;
  image: string;
}

interface Props {
  cart: Food[];
  onCheckout: () => void;
}

const Cart: React.FC<Props> = ({ cart, onCheckout }) => {
  const total = cart.reduce((sum, f) => sum + f.price, 0);

  return (
    <div className="fixed bottom-4 right-4 bg-white p-4 rounded-2xl shadow-lg w-72">
      <h2 className="font-bold text-lg">🛒 Cart</h2>
      {cart.length === 0 ? (
        <p className="text-gray-500">Cart is empty</p>
      ) : (
        <ul className="divide-y mt-2">
          {cart.map((item, i) => (
            <li key={i} className="flex justify-between py-1">
              <span>{item.name}</span>
              <span>฿{item.price}</span>
            </li>
          ))}
        </ul>
      )}
      <div className="mt-2 flex justify-between font-bold">
        <span>Total:</span>
        <span>฿{total}</span>
      </div>
      <button
        onClick={onCheckout}
        className="mt-3 w-full bg-yellow-400 py-2 rounded-lg font-semibold hover:bg-yellow-500"
      >
        Checkout
      </button>
    </div>
  );
};

export default Cart;
