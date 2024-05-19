<template>
    <div >
        <img :src="scsm_logo"></img>
    </div>
</template>


<script setup lang="ts">
import scsm_logo from "../assets/sarcasm_logo.jpeg"

const props = defineProps<{
    size: number
}>()
</script>

<style scoped>
img {
    width: v-bind(size + 'px');
    height: v-bind(size + 'px');
}
</style>