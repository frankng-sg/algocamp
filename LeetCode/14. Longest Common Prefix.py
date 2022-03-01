class Solution:
    def longestCommonPrefix(self, strPrefix):
        if not strPrefix:
            return ""
            
        strCommon = strPrefix[0]
        for i in range(len(strCommon)):
            for j in range(1, len(strPrefix)):
                try:
                    if strCommon[i] != strPrefix[j][i]:
                        if i == 0:
                            return ""
                        else:
                            return strCommon[:i]
                except:
                    if i == 0:
                        return ""
                    else:
                        return strCommon[:i]                    
        return strCommon
        