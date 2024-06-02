import React from "react";
import { defaults } from "chart.js/auto";
import { Line } from "react-chartjs-2";
import './../App.css'

defaults.maintainAspectRatio = false;
defaults.responsive = true;
defaults.plugins.title.display = true;
defaults.plugins.title.align = "start";
defaults.plugins.title.font.size = 20;
defaults.plugins.title.color = "black";

const speedtestData = [
  {
    id: 1,
    timestamp: '2024-06-02 21:33:36+00:00',
    isp: 'TIM',
    ping_latency: 1.777999997138977,
    ping_jitter: 0.24799999594688416,
    name_server: 'TIM Brasil',
    bandwidth_download: 38059738,
    bytes_download: 164567492,
    ping_download: 11.204000473022461,
    jitter_download: 2.694000005722046,
    bandwidth_upload: 23722236,
    bytes_upload: 144236595,
    ping_upload: 28.007999420166016,
    jitter_upload: 2.236999988555908,
  },
  {
    id: 2,
      timestamp: '2024-06-02 21:37:38+00:00',
      isp: 'TIM',
      ping_latency: 2.388000011444092,
      ping_jitter: 0.5189999938011169,
      name_server: 'TIM Brasil',
      bandwidth_download: 37596356,
      bytes_download: 174353232,
      ping_download: 12.795999526977539,
      jitter_download: 1.8799999952316284,
      bandwidth_upload: 24139130,
      bytes_upload: 177517644,
      ping_upload: 30.420000076293945,
      jitter_upload: 3.2090001106262207,
  },
  {
    id: 3,
      timestamp: '2024-06-02 21:41:41+00:00',
      isp: 'TIM',
      ping_latency: 1.9930000305175781,
      ping_jitter: 0.1459999978542328,
      name_server: 'TIM Brasil',
      bandwidth_download: 37461952,
      bytes_download: 265439404,
      ping_download: 17.8700008392334,
      jitter_download: 2.1730000972747803,
      bandwidth_upload: 24182807,
      bytes_upload: 196579348,
      ping_upload: 30.033000946044922,
      jitter_upload: 2.2160000801086426,
  },
  {
    id: 4,
      timestamp: '2024-06-02 21:45:35+00:00',
      isp: 'TIM',
      ping_latency: 1.9850000143051147,
      ping_jitter: 0.026000000536441803,
      name_server: 'TIM Brasil',
      bandwidth_download: 38006095,
      bytes_download: 246246052,
      ping_download: 15.866000175476074,
      jitter_download: 1.4290000200271606,
      bandwidth_upload: 24082083,
      bytes_upload: 86771173,
      ping_upload: 27.915000915527344,
      jitter_upload: 1.6579999923706055,
  },
  {
    id: 5,
      timestamp: '2024-06-02 21:49:37+00:00',
      isp: 'TIM',
      ping_latency: 1.968000054359436,
      ping_jitter: 0.39899998903274536,
      name_server: 'TIM Brasil',
      country_server: 'Brazil',
      bandwidth_download: 38214472,
      bytes_download: 165243416,
      ping_download: 11.668000221252441,
      jitter_download: 1.9190000295639038,
      bandwidth_upload: 24245772,
      bytes_upload: 173539344,
      ping_upload: 30.104999542236328,
      jitter_upload: 2.993000030517578,
  },
  {
    id: 6,
      timestamp: '2024-06-02 21:53:42+00:00',
      isp: 'TIM',
      ping_latency: 2.0859999656677246,
      ping_jitter: 0.3179999887943268,
      name_server: 'TIM Brasil',
      country_server: 'Brazil',
      bandwidth_download: 37971467,
      bytes_download: 268998960,
      ping_download: 17.749000549316406,
      jitter_download: 1.8519999980926514,
      bandwidth_upload: 24232914,
      bytes_upload: 169224602,
      ping_upload: 30.733999252319336,
      jitter_upload: 2.7899999618530273,
  }
]

export default function LineChart() {
  return (
    <div className="line-chart-card">
      <Line
        data={{
          labels: speedtestData.map((data) => data.timestamp),
          datasets: [
            {
              label: "Download in Mbps",
              data: speedtestData.map((data) => (data.bytes_download / 1048576)),
              backgroundColor: "rgba(6,79,240, 1.0)",
                borderColor: "rgba(6,79,240, 1.0)",
              },
              {
                label: "Upload in Mbps",
                data: speedtestData.map((data) => (data.bytes_upload / 1048576)),
                backgroundColor: "rgba(255, 48, 48, 0.5)",
                borderColor: "rgba(255, 48, 48, 1.0)",
                fill: true,
              },
            ],
          }}
          options={{
            elements: {
              line: {
                tension: 0.5,
              },
            },
            plugins: {
              title: {
                text: "Download and Upload Speeds",
              }
            },
          }}
        />
    </div>
  );
}
