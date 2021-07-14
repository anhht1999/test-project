<template>
    <MainSilder/>
    <MainProduct/>
</template>

<script>

import MainSilder from "./MainSilder.vue";
import MainProduct from "./MainProduct.vue";
export default {
    name:"HomePage",
    components: {
        MainSilder,
        MainProduct,
    }
}
</script>