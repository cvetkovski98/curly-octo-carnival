import React from "react";
import MachineList from "./MachineList";
import { ProductionLine } from "./types";

type ProductionLineItemProps = {
  line: ProductionLine;
};
const ProductionLineItem: React.FC<ProductionLineItemProps> = ({ line }) => {
  const { name, zone, factory } = line;
  return (
    <div className={"rounded-md bg-green-100 text-gray-600 mb-4"}>
      <div className="bg-green-300 px-4 py-5 border-b rounded-t-md border-gray-200 sm:px-6">
        <h3 className="text-lg leading-6 font-medium text-gray-800">{name}</h3>
      </div>
      <div className={"p-6"}>
        <p className="mb-4 font-bold">
          Location: {factory} @ {zone}
        </p>
        <MachineList machines={line.machines} />
      </div>
    </div>
  );
};

export default ProductionLineItem;
