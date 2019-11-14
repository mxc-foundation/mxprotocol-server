// https://react.i18next.com/getting-started
import i18next from "i18next";
import { initReactI18next } from "react-i18next";

import { name } from "../../package.json";
import Debug from "../util/debug";
import SessionStore from "../stores/SessionStore";
import { en, zhTW } from "./locales";

// Labels must match JSON filenames in locales directory
const supportedLanguages = [
  {
    label: "en",
    value: "English"
  },
  {
    label: "zhTW",
    value: "Chinese (Traditional)"
  }
];
let resourceEnglishNS = {};
let resourceChineseTraditionalNS = {};
resourceEnglishNS[name] = en;
resourceChineseTraditionalNS[name] = zhTW;
const packageNS = Object.keys(resourceEnglishNS)[0].toString();
const moduleNS = "i18n";
const menuNS = `${packageNS}-${moduleNS}`;
const debug = Debug(menuNS);
console.log('resourceEnglishNS: ', resourceEnglishNS);

const i18n = i18next;
i18n
  .use(initReactI18next)
  .init({
    debug: true,
    defaultNS: packageNS,
    fallbackLng: ["en-US", "en", "zhTW", "zhCN"],
    interpolation: {
      escapeValue: false
    },
    lng: SessionStore.getLanguageID() || "en",
    ns: [packageNS],
    // https://react.i18next.com/misc/using-with-icu-format
    react: {
      wait: true,
      bindI18n: "languageChanged loaded",
      bindStore: "added removed",
      nsMode: "default"
    },
    resources: {
      en: resourceEnglishNS,
      zhTW: resourceChineseTraditionalNS
    },
    saveMissing: true
  })
  .then(() => debug("success"))
  .catch(error => debug("failure", error));

const storedLanguage = SessionStore.getLanguageID();

i18next.on("initialized", options => {
  debug("Detected initialisation of i18n");
});

i18next.on("loaded", loaded => {
  debug("Detected success loading resources: ", loaded);
});

i18next.on("failedLoading", (lng, ns, msg) => {
  debug("Detected failure loading resources: ", lng, ns, msg);
});

// saveMissing must be configured to `true`
i18next.on("missingKey", (lngs, namespace, key, res) => {
  debug("Detected missing key: ", lngs, namespace, key, res);
});

i18next.store.on("added", (lng, ns) => {
  debug("Detected resources added: ", lng, ns);
});

i18next.store.on("removed", (lng, ns) => {
  debug("Detected resources removed: ", lng, ns);
});

// https://www.i18next.com/overview/api#changelanguage
i18next.on("languageChanged", lng => {
  debug("Detected language change to: ", lng);
});

export default i18n;

export { supportedLanguages, packageNS };
