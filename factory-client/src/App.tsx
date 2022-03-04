import React, { useState } from "react";
import instance from "./axiosInstance";
import ProductionLineItem from "./ProductionLineItem";
import { GetInventoryResponse } from "./types";

const App: React.FC = () => {
  const [inventoryResponse, setInventoryResponse] = useState<GetInventoryResponse>();

  const onButtonClick = async () => {
    instance
      .get<GetInventoryResponse>("/inventory")
      .then((resp) => setInventoryResponse(() => resp.data));
  };

  return (
    <div className={"p-6 m-6"}>
      <button
        type="button"
        onClick={onButtonClick}
        className="mb-4 inline-flex items-center px-6 py-3 border border-transparent text-base font-medium rounded-md shadow-sm text-white bg-indigo-600 hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
      >
        Receive inventory list
      </button>
        {inventoryResponse && <div className={"pb-4"}>Type: {inventoryResponse.type}</div>}
      {inventoryResponse?.lines.map((line, index) => (
        <ProductionLineItem line={line} key={index} />
      ))}
    </div>
  );
};

export default App;
