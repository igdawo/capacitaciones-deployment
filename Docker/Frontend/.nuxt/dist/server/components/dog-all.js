exports.ids = [1];
exports.modules = {

/***/ 37:
/***/ (function(module, exports, __webpack_require__) {

// style-loader: Adds some css to the DOM by adding a <style> tag

// load the styles
var content = __webpack_require__(42);
if(content.__esModule) content = content.default;
if(typeof content === 'string') content = [[module.i, content, '']];
if(content.locals) module.exports = content.locals;
// add CSS to SSR context
var add = __webpack_require__(4).default
module.exports.__inject__ = function (context) {
  add("7a974fbf", content, true, context)
};

/***/ }),

/***/ 41:
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony import */ var _node_modules_vue_style_loader_index_js_ref_3_oneOf_1_0_node_modules_css_loader_dist_cjs_js_ref_3_oneOf_1_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_ref_3_oneOf_1_2_node_modules_nuxt_components_dist_loader_js_ref_0_0_node_modules_vue_loader_lib_index_js_vue_loader_options_DogAll_vue_vue_type_style_index_0_id_67ae9bd6_prod_lang_css___WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(37);
/* harmony import */ var _node_modules_vue_style_loader_index_js_ref_3_oneOf_1_0_node_modules_css_loader_dist_cjs_js_ref_3_oneOf_1_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_ref_3_oneOf_1_2_node_modules_nuxt_components_dist_loader_js_ref_0_0_node_modules_vue_loader_lib_index_js_vue_loader_options_DogAll_vue_vue_type_style_index_0_id_67ae9bd6_prod_lang_css___WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(_node_modules_vue_style_loader_index_js_ref_3_oneOf_1_0_node_modules_css_loader_dist_cjs_js_ref_3_oneOf_1_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_ref_3_oneOf_1_2_node_modules_nuxt_components_dist_loader_js_ref_0_0_node_modules_vue_loader_lib_index_js_vue_loader_options_DogAll_vue_vue_type_style_index_0_id_67ae9bd6_prod_lang_css___WEBPACK_IMPORTED_MODULE_0__);
/* harmony reexport (unknown) */ for(var __WEBPACK_IMPORT_KEY__ in _node_modules_vue_style_loader_index_js_ref_3_oneOf_1_0_node_modules_css_loader_dist_cjs_js_ref_3_oneOf_1_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_ref_3_oneOf_1_2_node_modules_nuxt_components_dist_loader_js_ref_0_0_node_modules_vue_loader_lib_index_js_vue_loader_options_DogAll_vue_vue_type_style_index_0_id_67ae9bd6_prod_lang_css___WEBPACK_IMPORTED_MODULE_0__) if(["default"].indexOf(__WEBPACK_IMPORT_KEY__) < 0) (function(key) { __webpack_require__.d(__webpack_exports__, key, function() { return _node_modules_vue_style_loader_index_js_ref_3_oneOf_1_0_node_modules_css_loader_dist_cjs_js_ref_3_oneOf_1_1_node_modules_vue_loader_lib_loaders_stylePostLoader_js_node_modules_postcss_loader_dist_cjs_js_ref_3_oneOf_1_2_node_modules_nuxt_components_dist_loader_js_ref_0_0_node_modules_vue_loader_lib_index_js_vue_loader_options_DogAll_vue_vue_type_style_index_0_id_67ae9bd6_prod_lang_css___WEBPACK_IMPORTED_MODULE_0__[key]; }) }(__WEBPACK_IMPORT_KEY__));


/***/ }),

/***/ 42:
/***/ (function(module, exports, __webpack_require__) {

// Imports
var ___CSS_LOADER_API_IMPORT___ = __webpack_require__(3);
var ___CSS_LOADER_EXPORT___ = ___CSS_LOADER_API_IMPORT___(function(i){return i[1]});
// Module
___CSS_LOADER_EXPORT___.push([module.i, "#box-dog-All{height:450px;overflow-y:scroll}#button{background-color:#5f9ea0;color:#f5f5f5;font-size:100%;font-weight:700;width:50%}", ""]);
// Exports
___CSS_LOADER_EXPORT___.locals = {};
module.exports = ___CSS_LOADER_EXPORT___;


/***/ }),

/***/ 46:
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
// ESM COMPAT FLAG
__webpack_require__.r(__webpack_exports__);

// CONCATENATED MODULE: ./node_modules/babel-loader/lib??ref--2-0!./node_modules/vue-loader/lib/loaders/templateLoader.js??ref--6!./node_modules/@nuxt/components/dist/loader.js??ref--0-0!./node_modules/vue-loader/lib??vue-loader-options!./components/DogAll.vue?vue&type=template&id=67ae9bd6&
var render = function render() {
  var _vm = this,
    _c = _vm._self._c;
  return _c('div', {
    staticClass: "box container"
  }, [_vm._ssrNode("<div id=\"box-dog-All\"><table id=\"table-box\" class=\"table is-hoverable is-fullwidth\"><thead id=\"header-table\"><tr><th>Nombre</th> <th>Raza</th> <th>Edad</th> <th>Color</th> <th>Información</th></tr></thead> <tbody id=\"table-body\">" + _vm._ssrList(_vm.dogs, function (Dog, index) {
    return "<tr><td style=\"width: 20%\">" + _vm._ssrEscape(_vm._s(Dog.Nombre)) + "</td> <td style=\"width: 20%\">" + _vm._ssrEscape(_vm._s(Dog.Raza)) + "</td> <td style=\"width: 20%\">" + _vm._ssrEscape(_vm._s(Dog.Edad)) + "</td> <td style=\"width: 20%\">" + _vm._ssrEscape(_vm._s(Dog.Color)) + "</td> <td style=\"width: 20%; justify-content: center;align-items: center; \"><button id=\"button\">&gt;&gt;&gt;</button></td></tr>";
  }) + "</tbody></table></div>")]);
};
var staticRenderFns = [];

// CONCATENATED MODULE: ./components/DogAll.vue?vue&type=template&id=67ae9bd6&

// CONCATENATED MODULE: ./node_modules/babel-loader/lib??ref--2-0!./node_modules/@nuxt/components/dist/loader.js??ref--0-0!./node_modules/vue-loader/lib??vue-loader-options!./components/DogAll.vue?vue&type=script&lang=js&
/* harmony default export */ var DogAllvue_type_script_lang_js_ = ({
  data() {
    return {
      dogs: []
    };
  },
  methods: {
    async fetchAllDogs() {
      const path = "http://localhost:8000" + '/perro/';
      try {
        const res = await this.$axios.get(path);
        const data = res.data;
        this.dogs = data;
      } catch (error) {
        console.log(error);
      }
    },
    fetchDogData(dogID) {
      this.$router.push('/perros/' + dogID);
    }
  },
  created() {
    this.fetchAllDogs();
  }
});
// CONCATENATED MODULE: ./components/DogAll.vue?vue&type=script&lang=js&
 /* harmony default export */ var components_DogAllvue_type_script_lang_js_ = (DogAllvue_type_script_lang_js_); 
// EXTERNAL MODULE: ./node_modules/vue-loader/lib/runtime/componentNormalizer.js
var componentNormalizer = __webpack_require__(2);

// CONCATENATED MODULE: ./components/DogAll.vue



function injectStyles (context) {
  
  var style0 = __webpack_require__(41)
if (style0.__inject__) style0.__inject__(context)

}

/* normalize component */

var component = Object(componentNormalizer["a" /* default */])(
  components_DogAllvue_type_script_lang_js_,
  render,
  staticRenderFns,
  false,
  injectStyles,
  null,
  "76ec8035"
  
)

/* harmony default export */ var DogAll = __webpack_exports__["default"] = (component.exports);

/***/ })

};;
//# sourceMappingURL=dog-all.js.map