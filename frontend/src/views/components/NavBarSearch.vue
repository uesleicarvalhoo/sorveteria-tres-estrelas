<template>
  <control ref="root" placeholder="Search (ctrl+k)" ctrl-k-focus transparent borderless />
</template>

<script>
import Control from "./Control.vue"

export default {
  name: "NavBarSearch",
  components: {
    Control
  }
}
</script>
