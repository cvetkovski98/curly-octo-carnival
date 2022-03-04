import React from "react";
import { Machine } from "./types";
type MachineItemProps = {
  machines: Array<Machine>;
};

const dateFromTimestamp = (timestamp: number) => {
  return new Date(timestamp).toLocaleDateString();
};

const MachineList: React.FC<MachineItemProps> = ({ machines }) => {
  return (
    <div className={"grid grid-cols-12 gap-6"}>
      {machines.map(
        ({ id, name, serialNumber, installedAt, lastServicedAt }) => (
          <div className={"col-span-4 bg-red-100 rounded-xl"} key={id}>
            <div className="bg-red-300 px-4 py-5 border-b rounded-t-md sm:px-6">
              <h3 className="text-lg leading-6 font-medium text-gray-800">
                {name}
              </h3>
            </div>
            <div className="p-6">
              <div>
                <span>Serial number: </span>
                <span className="font-bold">{serialNumber}</span>
              </div>
              <div>
                <span>Installed at: </span>
                <span className="font-bold">
                  {dateFromTimestamp(installedAt.seconds)}
                </span>
              </div>
              <div>
                <span>Last service at: </span>
                <span className="font-bold">
                  {dateFromTimestamp(lastServicedAt.seconds)}
                </span>
              </div>
            </div>
          </div>
        )
      )}
    </div>
  );
};

export default MachineList;
