import axios from "axios";
import "./App.css";

import { GeoObject, Map, Placemark, YMaps } from "@pbe/react-yandex-maps";
import { useEffect, useState } from "react";
import uniqolor from "uniqolor";

interface Coordinates {
  Latitude: number;
  Longitude: number;
}

interface Orders {
  ID: number;
  CourierID: number;
  Start: Coordinates;
  End: Coordinates;
  Price: number;
}

interface Courier {
  ID: number;
  Coordinates: Coordinates;
  FullPathTime: number;
  Orders: Orders[];
  Path: number[];
  Penalty: number;
  Speed: number;
  Color?: string;
}

const App = () => {
  const [couriers, setCouriers] = useState<Courier[]>([]);

  const getCouriers = async () => {
    axios.get("http://127.0.0.1:8080/courier").then((res) => {
      const couriers: Courier[] = res.data;
      for (let i = 0; i < couriers.length; i++) {
        couriers[i].Color = uniqolor.random().color;
      }
      setCouriers(couriers);
    });
  };

  useEffect(() => {
    getCouriers();
  }, []);

  return (
    <YMaps>
      <Map
        style={{
          height: "100vh",
          width: "100vw",
        }}
        defaultState={{ center: [62.02722, 129.732178], zoom: 14 }}
      >
        {couriers.map((courier) => (
          <Placemark
            options={{
              iconColor: courier.Color,
            }}
            onClick={() => alert(courier.Path)}
            geometry={[
              courier.Coordinates.Latitude,
              courier.Coordinates.Longitude,
            ]}
          />
        ))}
        {couriers.map((courier) =>
          courier.Orders.map((order) => (
            <GeoObject
              geometry={{
                type: "LineString",
                coordinates: [
                  [order.Start.Latitude, order.Start.Longitude],
                  [order.End.Latitude, order.End.Longitude],
                ],
              }}
              options={{
                geodesic: true,
                strokeWidth: 5,
                strokeColor: courier.Color,
              }}
            />
          )),
        )}
      </Map>
    </YMaps>
  );
};

export default App;
