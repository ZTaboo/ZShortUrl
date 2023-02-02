<template>
    <div class="header-full">
        <div class="box header">
            <div>
                <img class="logo" src="@/assets/icon.png" alt="">
                <h3>ZShortUrl</h3>
            </div>
            <div>
                <n-button type="info" size="small">登录</n-button>
                <n-button size="small" color="white" text-color="black" strong>注册</n-button>
            </div>
        </div>
    </div>
</template>

<script lang="ts" setup>

</script>

<style scoped>
.logo {
    height: 60px;
}

.header {
    margin: auto;
    height: 100%;
    display: flex;
    justify-content: space-between;
}

.header div {
    display: flex;
    align-items: center;
    user-select: none;
}

.header > div:last-child > button {
    margin-left: 10px;
}

.header > div:last-child > button:last-child{
    border: 1px solid #ccc;
}

.header-full {
    width: 100%;
    height: 60px;
}
</style>