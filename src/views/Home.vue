<template>
    <Header></Header>
    <div class="margin content-full">

        <div class="box content">
            <h2>简单易用的渠道短链接统计工具</h2>
            <div class="con-bg">
                <div class="con">
                    <n-input style="margin-left: 10px" status="success" v-model:value="content"></n-input>
                    <n-button type="info" style="margin: 0 10px 0 10px;" @click="btn">生成短连接</n-button>
                </div>
            </div>
            <div class="bg-box">
                <object :data="bg" class="bg-box" type="image/svg+xml"/>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import Header from "../components/public/Header.vue";
import bg from "../assets/svg/home-bg.svg"
import axios from "axios";
import {h, ref} from "vue";
import {useDialog} from 'naive-ui'

const content = ref('')
const notification = useMessage()
const dialog = useDialog()
const btn = () => {
    axios.post("http://localhost:8080/genNewUrl", {url: content.value}).then(r => {
        if (r.data.code === 200) {
            dialog.success({
                title: '生成完成',
                content: () => h("a", {href: r.data.msg}, r.data.msg),
            })
        } else {
            notification.error(`生成失败：${r.data.msg}`)
        }
    })
}

const con = defineProps({})
</script>

<style scoped>
@import "@/assets/css/home.css";
</style>