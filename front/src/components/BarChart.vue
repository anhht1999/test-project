<template>
  <div class="example">
    <apexcharts width="800" height="450" type="bar" :options="chartOptions" :series="series"></apexcharts>
  </div>
</template>

<script>
import VueApexCharts from 'vue3-apexcharts';

export default {
  name: 'BarChart',
  components: {
    apexcharts: VueApexCharts,
  },
  props: {
    date: Array,
    rev: Array,
  },
  data() {
    return {
      chartOptions: {
        chart: {
          id: 'date'
        },
        xaxis: {
          categories: this.date
        }
      },
      series: [{
        name: 'rev',
        data: this.rev
      }]
    }
  },
  
}
</script>
