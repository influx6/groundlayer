export namespace Objects {
  export function PatchWith(elem: any, attrName: string, attrs: any): void {
    elem[attrName] = attrs;
  }

  export function GetAttrWith(elem: any, attrName: string): any {
    const value = elem[attrName];
    return value ? value : '';
  }

  export function isNullOrUndefined(elem: any): boolean {
    return elem === null || elem === undefined;
  }

  export function isAny(elem: any, ...values: any): boolean {
    for (let index of values) {
      if (elem === index) {
        return true;
      }
    }
    return false;
  }
}
