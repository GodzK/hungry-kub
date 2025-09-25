import React from "react";

interface FoodCardProps {
  food: {
    name: string;
    price: number;
    image: string;
  };
  onAdd: (food: any) => void;
  className?: string;
}

const FoodCard: React.FC<FoodCardProps> = ({ food, onAdd, className }) => {
  return (
    <div className={`bg-white rounded-lg overflow-hidden shadow-lg ${className}`}>
      <img
        src={food.image}
        alt={food.name}
        className="w-full h-48 object-cover"
      />
      <div className="p-4">
        <h2 className="text-lg font-bold">{food.name}</h2>
        <p className="text-gray-500">฿{food.price}</p>
        <button
          className="mt-3 w-full bg-primary text-white py-2 rounded-lg hover:bg-green-600 transition-colors"
          onClick={() => onAdd(food)}
        >
          เพิ่มไปยังรถเข็น
        </button>
      </div>
    </div>
  );
};

export default FoodCard;
