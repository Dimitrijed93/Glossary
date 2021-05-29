<template>
  <div class="content">
    <div class="container-fluid">
      <div class="col-6 d-flex justify-content-start">
        <div class="margin-items d-flex flex-grow-1">
          <select v-model="folders.folderId" class="form-control">
            <option v-for="folder in folders.folderOptions" v-bind:value="folder.id">
              {{ itemValue(folder) }}
            </option>
          </select>
        </div>
        <div class="margin-items d-flex flex-grow-4">
          <button
            id="newFolder"
            class="btn btn-fill btn-primary"
            @click="newFolder"
          >
            New Folder
          </button>
        </div>
        <div class="margin-items d-flex flex-grow-2">
          <button
            id="editFolder"
            v-if="folders.folderId"
            class="btn btn-fill btn-primary"
            @click="edit"
          >
            Edit Folder
          </button>
        </div>
      </div>
      <folder-modal
        v-on:getAll="getAll"
        v-bind:languageOptions="folders.languageOptions"
        ref="child"
      ></folder-modal>
    </div>
  </div>
</template>

<<script>
import LTable from "src/components/Table.vue";
import FolderModal from "./FolderModal.vue";
import Card from "src/components/Cards/Card.vue";
import axios from "axios";
import URL from 'src/util/url';
export default {
  components: {
    LTable,
    FolderModal,
    Card,
  },
  data() {
    return {
      folders: {
        folderId: 0,
        folderOptions: [],
        languageOptions: [],
        folderName: "",
        source: "",
        destination: "",
        BASE_URL: URL.LOCAL_BASE
      },
    };
  },
  mounted() {
    this.getAll();
  },
  methods: {
    itemValue(folder) {
        return folder["name"];
    },
    newFolder() {
      this.$refs.child.newFolder();
    },
    getAll() {
      axios.get(`${this.folders.BASE_URL}/folders-and-languages`).then((response) => {
        this.folders["folderOptions"] = response.data.folders;
        this.folders["languageOptions"] = response.data.languages;

      });
    },
    edit(ev, item) {
      let id = this.folders.folderId;
      this.$refs.child.showEditModal(id);

    },
    deleteItem(item) {
      axios.delete(`${this.folders.BASE_URL}/types/${item.id}`).then(() => {
        this.getAll();
      });
    },
  },
};
</script>
<style scoped>
.margin-items {
  margin: 5px;
}
</style>