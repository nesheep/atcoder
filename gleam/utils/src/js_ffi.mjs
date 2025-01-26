import { readFileSync } from "node:fs";

export function read_all() {
  return readFileSync(0, "utf8");
}
