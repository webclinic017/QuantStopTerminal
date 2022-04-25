<template>
  <Layout title="ChartTypeKLineChart">
    <div id="chart-type-k-line" class="k-line-chart"/>
    <div
      class="k-line-chart-menu-container">
      <button
        v-for="({ key, text }) in chartTypes"
        :key="key"
        v-on:click="setChartType(key)">
        {{text}}
      </button>
    </div>
  </Layout>
</template>

<script>
import { init, dispose } from 'klinecharts'
import generatedKLineDataList from "./generatedKLineDataList";
import Layout from "./Layout"

function getLanguageOptions (language) {
  return {
    candle: {
      tooltip: {
        labels: language === 'zh-CN'
            ? ['时间：', '开：', '收：', '高：', '低：', '成交量：']
            : language === 'zh-HK'
                ? ['時間：', '開：', '收：', '高：', '低：', '成交量：']
                : ['T: ', 'O: ', 'C: ', 'H: ', 'L: ', 'V: ']
      }
    }
  }
}

export default {
  name: 'ChartTypeKLineChart',
  components: {Layout},
  data: function () {
    return {
      chartTypes: [
        { key: 'candle_solid', text: 'candle solid' },
        { key: 'candle_stroke', text: 'candle stroke' },
        { key: 'candle_up_stroke', text: 'candle up stroke' },
        { key: 'candle_down_stroke', text: 'candle down stroke' },
        { key: 'ohlc', text: 'OHLC' },
        { key: 'area', text: 'area' }
      ]
    }
  },
  mounted: function () {
    this.kLineChart = init('chart-type-k-line')
    this.kLineChart.setStyleOptions(getLanguageOptions("english"))
    this.kLineChart.applyNewData(generatedKLineDataList())
  },
  methods: {
    setChartType: function (type) {
      this.kLineChart.setStyleOptions({
        candle: {
          type
        }
      })
    },

  },
  destroyed: function () {
    dispose('chart-type-k-line')
  }
}
</script>
