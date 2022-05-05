// Code generated by "stringer -type=FeatureID,Vendor"; DO NOT EDIT.

package cpuid

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ADX-1]
	_ = x[AESNI-2]
	_ = x[AMD3DNOW-3]
	_ = x[AMD3DNOWEXT-4]
	_ = x[AMXBF16-5]
	_ = x[AMXINT8-6]
	_ = x[AMXTILE-7]
	_ = x[AVX-8]
	_ = x[AVX2-9]
	_ = x[AVX512BF16-10]
	_ = x[AVX512BITALG-11]
	_ = x[AVX512BW-12]
	_ = x[AVX512CD-13]
	_ = x[AVX512DQ-14]
	_ = x[AVX512ER-15]
	_ = x[AVX512F-16]
	_ = x[AVX512FP16-17]
	_ = x[AVX512IFMA-18]
	_ = x[AVX512PF-19]
	_ = x[AVX512VBMI-20]
	_ = x[AVX512VBMI2-21]
	_ = x[AVX512VL-22]
	_ = x[AVX512VNNI-23]
	_ = x[AVX512VP2INTERSECT-24]
	_ = x[AVX512VPOPCNTDQ-25]
	_ = x[AVXSLOW-26]
	_ = x[BMI1-27]
	_ = x[BMI2-28]
	_ = x[CETIBT-29]
	_ = x[CETSS-30]
	_ = x[CLDEMOTE-31]
	_ = x[CLMUL-32]
	_ = x[CLZERO-33]
	_ = x[CMOV-34]
	_ = x[CMPXCHG8-35]
	_ = x[CPBOOST-36]
	_ = x[CX16-37]
	_ = x[ENQCMD-38]
	_ = x[ERMS-39]
	_ = x[F16C-40]
	_ = x[FMA3-41]
	_ = x[FMA4-42]
	_ = x[FXSR-43]
	_ = x[FXSROPT-44]
	_ = x[GFNI-45]
	_ = x[HLE-46]
	_ = x[HTT-47]
	_ = x[HWA-48]
	_ = x[HYPERVISOR-49]
	_ = x[IBPB-50]
	_ = x[IBS-51]
	_ = x[IBSBRNTRGT-52]
	_ = x[IBSFETCHSAM-53]
	_ = x[IBSFFV-54]
	_ = x[IBSOPCNT-55]
	_ = x[IBSOPCNTEXT-56]
	_ = x[IBSOPSAM-57]
	_ = x[IBSRDWROPCNT-58]
	_ = x[IBSRIPINVALIDCHK-59]
	_ = x[INT_WBINVD-60]
	_ = x[INVLPGB-61]
	_ = x[LAHF-62]
	_ = x[LZCNT-63]
	_ = x[MCAOVERFLOW-64]
	_ = x[MCOMMIT-65]
	_ = x[MMX-66]
	_ = x[MMXEXT-67]
	_ = x[MOVBE-68]
	_ = x[MOVDIR64B-69]
	_ = x[MOVDIRI-70]
	_ = x[MPX-71]
	_ = x[MSRIRC-72]
	_ = x[NX-73]
	_ = x[OSXSAVE-74]
	_ = x[POPCNT-75]
	_ = x[RDPRU-76]
	_ = x[RDRAND-77]
	_ = x[RDSEED-78]
	_ = x[RDTSCP-79]
	_ = x[RTM-80]
	_ = x[RTM_ALWAYS_ABORT-81]
	_ = x[SCE-82]
	_ = x[SERIALIZE-83]
	_ = x[SGX-84]
	_ = x[SGXLC-85]
	_ = x[SHA-86]
	_ = x[SSE-87]
	_ = x[SSE2-88]
	_ = x[SSE3-89]
	_ = x[SSE4-90]
	_ = x[SSE42-91]
	_ = x[SSE4A-92]
	_ = x[SSSE3-93]
	_ = x[STIBP-94]
	_ = x[SUCCOR-95]
	_ = x[TBM-96]
	_ = x[TSXLDTRK-97]
	_ = x[VAES-98]
	_ = x[VMX-99]
	_ = x[VPCLMULQDQ-100]
	_ = x[WAITPKG-101]
	_ = x[WBNOINVD-102]
	_ = x[X87-103]
	_ = x[XOP-104]
	_ = x[XGETBV1-105]
	_ = x[XSAVE-106]
	_ = x[XSAVEC-107]
	_ = x[XSAVEOPT-108]
	_ = x[XSAVES-109]
	_ = x[AESARM-110]
	_ = x[ARMCPUID-111]
	_ = x[ASIMD-112]
	_ = x[ASIMDDP-113]
	_ = x[ASIMDHP-114]
	_ = x[ASIMDRDM-115]
	_ = x[ATOMICS-116]
	_ = x[CRC32-117]
	_ = x[DCPOP-118]
	_ = x[EVTSTRM-119]
	_ = x[FCMA-120]
	_ = x[FP-121]
	_ = x[FPHP-122]
	_ = x[GPA-123]
	_ = x[JSCVT-124]
	_ = x[LRCPC-125]
	_ = x[PMULL-126]
	_ = x[SHA1-127]
	_ = x[SHA2-128]
	_ = x[SHA3-129]
	_ = x[SHA512-130]
	_ = x[SM3-131]
	_ = x[SM4-132]
	_ = x[SVE-133]
	_ = x[lastID-134]
	_ = x[firstID-0]
}

const _FeatureID_name = "firstIDADXAESNIAMD3DNOWAMD3DNOWEXTAMXBF16AMXINT8AMXTILEAVXAVX2AVX512BF16AVX512BITALGAVX512BWAVX512CDAVX512DQAVX512ERAVX512FAVX512FP16AVX512IFMAAVX512PFAVX512VBMIAVX512VBMI2AVX512VLAVX512VNNIAVX512VP2INTERSECTAVX512VPOPCNTDQAVXSLOWBMI1BMI2CETIBTCETSSCLDEMOTECLMULCLZEROCMOVCMPXCHG8CPBOOSTCX16ENQCMDERMSF16CFMA3FMA4FXSRFXSROPTGFNIHLEHTTHWAHYPERVISORIBPBIBSIBSBRNTRGTIBSFETCHSAMIBSFFVIBSOPCNTIBSOPCNTEXTIBSOPSAMIBSRDWROPCNTIBSRIPINVALIDCHKINT_WBINVDINVLPGBLAHFLZCNTMCAOVERFLOWMCOMMITMMXMMXEXTMOVBEMOVDIR64BMOVDIRIMPXMSRIRCNXOSXSAVEPOPCNTRDPRURDRANDRDSEEDRDTSCPRTMRTM_ALWAYS_ABORTSCESERIALIZESGXSGXLCSHASSESSE2SSE3SSE4SSE42SSE4ASSSE3STIBPSUCCORTBMTSXLDTRKVAESVMXVPCLMULQDQWAITPKGWBNOINVDX87XOPXGETBV1XSAVEXSAVECXSAVEOPTXSAVESAESARMARMCPUIDASIMDASIMDDPASIMDHPASIMDRDMATOMICSCRC32DCPOPEVTSTRMFCMAFPFPHPGPAJSCVTLRCPCPMULLSHA1SHA2SHA3SHA512SM3SM4SVElastID"

var _FeatureID_index = [...]uint16{0, 7, 10, 15, 23, 34, 41, 48, 55, 58, 62, 72, 84, 92, 100, 108, 116, 123, 133, 143, 151, 161, 172, 180, 190, 208, 223, 230, 234, 238, 244, 249, 257, 262, 268, 272, 280, 287, 291, 297, 301, 305, 309, 313, 317, 324, 328, 331, 334, 337, 347, 351, 354, 364, 375, 381, 389, 400, 408, 420, 436, 446, 453, 457, 462, 473, 480, 483, 489, 494, 503, 510, 513, 519, 521, 528, 534, 539, 545, 551, 557, 560, 576, 579, 588, 591, 596, 599, 602, 606, 610, 614, 619, 624, 629, 634, 640, 643, 651, 655, 658, 668, 675, 683, 686, 689, 696, 701, 707, 715, 721, 727, 735, 740, 747, 754, 762, 769, 774, 779, 786, 790, 792, 796, 799, 804, 809, 814, 818, 822, 826, 832, 835, 838, 841, 847}

func (i FeatureID) String() string {
	if i < 0 || i >= FeatureID(len(_FeatureID_index)-1) {
		return "FeatureID(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _FeatureID_name[_FeatureID_index[i]:_FeatureID_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[VendorUnknown-0]
	_ = x[Intel-1]
	_ = x[AMD-2]
	_ = x[VIA-3]
	_ = x[Transmeta-4]
	_ = x[NSC-5]
	_ = x[KVM-6]
	_ = x[MSVM-7]
	_ = x[VMware-8]
	_ = x[XenHVM-9]
	_ = x[Bhyve-10]
	_ = x[Hygon-11]
	_ = x[SiS-12]
	_ = x[RDC-13]
	_ = x[Ampere-14]
	_ = x[ARM-15]
	_ = x[Broadcom-16]
	_ = x[Cavium-17]
	_ = x[DEC-18]
	_ = x[Fujitsu-19]
	_ = x[Infineon-20]
	_ = x[Motorola-21]
	_ = x[NVIDIA-22]
	_ = x[AMCC-23]
	_ = x[Qualcomm-24]
	_ = x[Marvell-25]
	_ = x[lastVendor-26]
}

const _Vendor_name = "VendorUnknownIntelAMDVIATransmetaNSCKVMMSVMVMwareXenHVMBhyveHygonSiSRDCAmpereARMBroadcomCaviumDECFujitsuInfineonMotorolaNVIDIAAMCCQualcommMarvelllastVendor"

var _Vendor_index = [...]uint8{0, 13, 18, 21, 24, 33, 36, 39, 43, 49, 55, 60, 65, 68, 71, 77, 80, 88, 94, 97, 104, 112, 120, 126, 130, 138, 145, 155}

func (i Vendor) String() string {
	if i < 0 || i >= Vendor(len(_Vendor_index)-1) {
		return "Vendor(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Vendor_name[_Vendor_index[i]:_Vendor_index[i+1]]
}
