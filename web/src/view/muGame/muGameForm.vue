<template>
  <div>
    <el-form :model="formData" label-position="right" label-width="80px">
      <el-form-item label="游戏标题:">
    <el-input v-model="formData.title" clearable placeholder="请输入" />
    </el-form-item>
    
      <el-form-item label="游戏logo:">
    <el-input v-model="formData.images" clearable placeholder="请输入" />
    </el-form-item>
    
      <el-form-item label="介绍内容:">
    <el-input v-model="formData.content" clearable placeholder="请输入" />
    </el-form-item>
    
      <el-form-item label="游戏连接:">
    <el-input v-model="formData.url" clearable placeholder="请输入" />
    </el-form-item>
    
      <el-form-item label="游戏分类:对应mu_nav的id:">
    <el-switch active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" v-model="formData.type" clearable ></el-switch>
    </el-form-item>
    
      <el-form-item label="添加时间:">
    
      <el-date-picker type="date" placeholder="选择日期" v-model="formData.createTime" clearable></el-date-picker>
    </el-form-item>
    
      <el-form-item label="游戏连接打开方式：1.app打开；2.浏览器打开:">
    <el-input v-model.number="formData.gameopentype" clearable placeholder="请输入"/>
    </el-form-item>
    <el-form-item>
        <el-button size="mini" type="primary" @click="save">保存</el-button>
        <el-button size="mini" type="primary" @click="back">返回</el-button>
      </el-form-item>
    </el-form>
</div>
</template>

<script>
import {
  createMuGame,
  updateMuGame,
  findMuGame
} from '@/api/muGame' //  此处请自行替换地址
import infoList from '@/mixins/infoList'
export default {
  name: 'MuGame',
  mixins: [infoList],
  data() {
    return {
      type: '',
      formData: {
            title: '',
            images: '',
            content: '',
            url: '',
            type: false,
            createTime: new Date(),
            gameopentype: 0,
            
      }
    }
  },
  async created() {
    // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (this.$route.query.id) {
      const res = await findMuGame({ ID: this.$route.query.id })
      if (res.code === 0) {
        this.formData = res.data.remuGame
        this.type = 'update'
      }
    } else {
      this.type = 'create'
    }
    
  },
  methods: {
    async save() {
      let res
      switch (this.type) {
        case 'create':
          res = await createMuGame(this.formData)
          break
        case 'update':
          res = await updateMuGame(this.formData)
          break
        default:
          res = await createMuGame(this.formData)
          break
      }
      if (res.code === 0) {
        this.$message({
          type: 'success',
          message: '创建/更改成功'
        })
      }
    },
    back() {
      this.$router.go(-1)
    }
  }
}
</script>

<style>
</style>