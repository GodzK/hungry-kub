import React, { useEffect, useState } from "react";
import Navbar from "./components/Navbar";
import FoodCard from "./components/FoodCard";
import Cart from "./components/Cart";

interface Food {
  _id: string;
  name: string;
  price: number;
  image: string;
}

function App() {
  const [foods, setFoods] = useState<Food[]>([]);
  const [cart, setCart] = useState<Food[]>([]);

  useEffect(() => {
    fetch("http://localhost:8080/foods")
      .then((res) => res.json())
      .then((data) => setFoods(data));
  }, []);

  const addToCart = (food: Food) => {
    setCart([...cart, food]);
  };

  const checkout = async () => {
    const res = await fetch("http://localhost:8080/order", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(cart),
    });
    const data = await res.json();
    alert("✅ Order placed! ID: " + data.orderId);
    setCart([]);
  };

  return (
    <div className="min-h-screen bg-secondary">
      <Navbar cartCount={cart.length} />
      
      {/* Hero Section */}
      <div className="bg-primary text-white py-10">
        <div className="container mx-auto px-4 text-center">
          <h1 className="text-4xl font-bold mb-2">Welcome to Foodie</h1>
          <p className="text-lg">สั่งอาหารได้ง่ายและรวดเร็ว เหมือน LINE MAN</p>
        </div>
      </div>

      {/* Food Grid */}
      <div className="container mx-auto px-4 py-6 grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-8">
        {foods.map((food) => (
          <FoodCard
            key={food._id}
            food={food}
            onAdd={addToCart}
            className="transition-transform transform hover:scale-105 shadow-card rounded-lg"
          />
        ))}
      </div>

      {/* Cart */}
      <Cart cart={cart} onCheckout={checkout} />
    </div>
  );
}

export default App;
