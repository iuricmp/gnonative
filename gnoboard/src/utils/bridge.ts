import { GoBridge } from "@gno/native_modules";

export const initBridge = async (): Promise<boolean> => {
  try {
    console.log("bridge methods: ", Object.keys(GoBridge));

    await GoBridge.initBridge();

    return true;
  } catch (err: any) {
    if (err?.message?.indexOf("already instantiated") !== -1) {
      console.log("bridge already started: ", err);
      return true;
    } else {
      console.error("unable to init bridge: ", err);
    }

    return false;
  }
};

export const closeBridge = async (): Promise<boolean> => {
  try {
    await GoBridge.closeBridge();
    return true;
  } catch (err: any) {
    console.error("unable to close bridge: ", err);
    return false;
  }
};

export const setPassword = async (password: string): Promise<boolean> => {
  await GoBridge.setPassword(password);
  return true;
};

export const generateRecoveryPhrase = async (): Promise<string> => {
  return await GoBridge.generateRecoveryPhrase();
};

export const listKeyInfo = async (): Promise<string[]> => {
  return await GoBridge.listKeyInfo();
};

export const createAccount = async (
  nameOrBech32: string,
  mnemonic: string,
  bip39Passw: string,
  password: string,
  account: Number,
  index: Number,
): Promise<string> => {
  return await GoBridge.createAccount(
    nameOrBech32,
    mnemonic,
    bip39Passw,
    password,
    account,
    index,
  );
};

export const selectAccount = async (nameOrBech32: string): Promise<string> => {
  return await GoBridge.selectAccount(nameOrBech32);
};
