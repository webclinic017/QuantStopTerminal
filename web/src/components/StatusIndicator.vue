<template>
  <div class="led-box d-inline-block">
    <div class="led-light d-inline-block" :class="connected ? 'led-green' : 'led-red'"></div>
  </div>
</template>

<script>
export default {
  name: 'StatusIndicator',
  props: {
    on: false
  },
  computed: {
    connected() {
      return this.on
    }
  }
};
</script>

<style scoped>
.led-box {
  padding-left: 10px;
  height: 24px;
  width: 24px;
}
.led-light {
  width: 24px;
  height: 24px;
  border-radius: 50%;
}
.led-red { background-color: #f00; box-shadow: rgba(0, 0, 0, 0.2) 0 -1px 7px 1px, inset #441313 0 -1px 9px, rgba(255, 0, 0, 0.5) 0 2px 12px; }
.led-green { background-color: #abff00; box-shadow: rgba(0, 0, 0, 0.2) 0 -1px 7px 1px, inset #304701 0 -1px 9px, #89ff00 0 2px 12px; }
</style>