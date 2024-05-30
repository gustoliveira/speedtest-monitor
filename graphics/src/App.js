import Chart from "chart.js/auto";
import { CategoryScale } from "chart.js";
import { useState } from "react";
import LineChart from "./components/line_chart";
import './App.css';

export const Data = [
  {
    id: 1,
    year: 2016,
    userGain: 80000,
    userLost: 823
  },
  {
    id: 2,
    year: 2017,
    userGain: 45677,
    userLost: 345
  },
  {
    id: 3,
    year: 2018,
    userGain: 78888,
    userLost: 555
  },
  {
    id: 4,
    year: 2019,
    userGain: 90000,
    userLost: 4555
  },
  {
    id: 5,
    year: 2020,
    userGain: 4300,
    userLost: 234
  }
];

Chart.register(CategoryScale);

export default function App() {
  const [chartData, _] = useState({
    labels: Data.map((data) => data.year), 
    datasets: [
      {
        label: "Users Gained ",
        data: Data.map((data) => data.userGain),
        borderColor: "black",
        borderWidth: 2
      },
    ]
  });

  return (
    <div className="app">
      <div className="chart-body">
        <LineChart chartData={chartData} />
      </div>
    </div>
  );
}

