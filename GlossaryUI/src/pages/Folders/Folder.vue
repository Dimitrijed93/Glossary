<template>
  <div class="content">
    <div class="container-fluid">
      <card>
      <template slot="header">
        <div class="row">
          <div class="col-12">
            <h4 class="card-title">Folders</h4>
          </div>
        </div>
      </template>
      <div class="col-6 d-flex justify-content-start">
        <div class="margin-items d-flex flex-grow-1">
          <select
            v-on:click="onSelectedFolder($event)"
            v-model="folders.folderId"
            class="form-control"
          >
            <option
              v-for="folder in folders.folderOptions"
              v-bind:value="folder.id"
            >
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
      <entry-list ref="entryList"> </entry-list>
      </card>
    </div>
  </div>
</template>

<<script>
import LTable from "src/components/Table.vue";
import FolderModal from "./FolderModal.vue";
import EntryList from "./EntryList.vue";
import Card from "src/components/Cards/Card.vue";
import axios from "axios";
import URL from 'src/util/url';
export default {
  components: {
    LTable,
    FolderModal,
    EntryList,
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
    onSelectedFolder(e) {
      if (e) {
        let id = e.target.value;
        if (id) {
          this.$refs.entryList.load(id);
        }
      }

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