<script setup>
import {
  UserOutlined,
  CreditCardOutlined,
  KeyOutlined
} from '@ant-design/icons-vue';
import { ref, reactive } from 'vue';
import { useRouter } from 'vue-router'

const router = useRouter()
const selectedKeys = ref(['1.1'])
const openKeys = ref(['1'])
const onClick = ({ item, key, keyPath }) => {
  const routerList = {
    1.1: "/account",
    2.1: "/client",
    3.1: "/token",
  }
  router.push(routerList[key])
}
</script>

<template>
  <a-layout class="layout">
    <a-layout-sider>
      <div class="logo" />
      <a-menu v-model:selectedKeys="selectedKeys" :open-keys="openKeys" @click="onClick" theme="dark" mode="inline">
        <a-sub-menu key="1">
          <template #title>
            <span>
              <user-outlined />
              <span>User</span>
            </span>
          </template>
          <a-menu-item key="1.1">用户列表</a-menu-item>
        </a-sub-menu>

        <a-sub-menu key="2">
          <template #title>
            <span>
              <credit-card-outlined />
              <span>Client</span>
            </span>
          </template>
          <a-menu-item key="2.1">Client列表</a-menu-item>
        </a-sub-menu>

        <a-sub-menu key="3">
          <template #title>
            <span>
              <key-outlined />
              <span>Token</span>
            </span>
          </template>
          <a-menu-item key="3.1">Token列表</a-menu-item>
        </a-sub-menu>

      </a-menu>
    </a-layout-sider>
    <a-layout>
      <a-layout-content :style="{ margin: '24px 16px 0' }">
        <div class="layout-content">
          <router-view></router-view>
        </div>
      </a-layout-content>
    </a-layout>
  </a-layout>
</template>

<style lang="less">
#app {
  // width: 100vw;
  height: 100vh;
  margin: 0;
  padding: 0;
  overflow: hidden;

  .layout {
    min-height: 100vh;
  }

  .layout-content {
    margin: 24px 16px 0;
    // height: 100%;
    padding: 24px;
    background: #fff;
  }
}
</style>
