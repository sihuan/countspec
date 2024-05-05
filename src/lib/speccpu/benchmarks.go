package speccpu

// 500.perlbench_r 	600.perlbench_s 	C 	362 	Perl interpreter
// 502.gcc_r 	602.gcc_s 	C 	1,304 	GNU C compiler
// 505.mcf_r 	605.mcf_s 	C 	3 	Route planning
// 520.omnetpp_r 	620.omnetpp_s 	C++ 	134 	Discrete Event simulation - computer network
// 523.xalancbmk_r 	623.xalancbmk_s 	C++ 	520 	XML to HTML conversion via XSLT
// 525.x264_r 	625.x264_s 	C 	96 	Video compression
// 531.deepsjeng_r 	631.deepsjeng_s 	C++ 	10 	Artificial Intelligence: alpha-beta tree search (Chess)
// 541.leela_r 	641.leela_s 	C++ 	21 	Artificial Intelligence: Monte Carlo tree search (Go)
// 548.exchange2_r 	648.exchange2_s 	Fortran 	1 	Artificial Intelligence: recursive solution generator (Sudoku)
// 557.xz_r 	657.xz_s 	C 	33 	General data compression

var INTRATE = []string{
	"500.perlbench_r",
	"502.gcc_r",
	"505.mcf_r",
	"520.omnetpp_r",
	"523.xalancbmk_r",
	"525.x264_r",
	"531.deepsjeng_r",
	"541.leela_r",
	"548.exchange2_r",
	"557.xz_r",
}

var INTSPEED = []string{
	"600.perlbench_s",
	"602.gcc_s",
	"605.mcf_s",
	"620.omnetpp_s",
	"623.xalancbmk_s",
	"625.x264_s",
	"631.deepsjeng_s",
	"641.leela_s",
	"648.exchange2_s",
	"657.xz_s",
}

// 503.bwaves_r 	603.bwaves_s 	Fortran 	1 	Explosion modeling
// 507.cactuBSSN_r 	607.cactuBSSN_s 	C++, C, Fortran 	257 	Physics: relativity
// 508.namd_r 	  	C++ 	8 	Molecular dynamics
// 510.parest_r 	  	C++ 	427 	Biomedical imaging: optical tomography with finite elements
// 511.povray_r 	  	C++, C 	170 	Ray tracing
// 519.lbm_r 	619.lbm_s 	C 	1 	Fluid dynamics
// 521.wrf_r 	621.wrf_s 	Fortran, C 	991 	Weather forecasting
// 526.blender_r 	  	C++, C 	1,577 	3D rendering and animation
// 527.cam4_r 	627.cam4_s 	Fortran, C 	407 	Atmosphere modeling
//   	628.pop2_s 	Fortran, C 	338 	Wide-scale ocean modeling (climate level)
// 538.imagick_r 	638.imagick_s 	C 	259 	Image manipulation
// 544.nab_r 	644.nab_s 	C 	24 	Molecular dynamics
// 549.fotonik3d_r 	649.fotonik3d_s 	Fortran 	14 	Computational Electromagnetics
// 554.roms_r 	654.roms_s 	Fortran 	210 	Regional ocean modeling

var FPRATE = []string{
	"503.bwaves_r",
	"507.cactuBSSN_r",
	"508.namd_r",
	"510.parest_r",
	"511.povray_r",
	"519.lbm_r",
	"521.wrf_r",
	"526.blender_r",
	"527.cam4_r",
	"538.imagick_r",
	"544.nab_r",
	"549.fotonik3d_r",
	"554.roms_r",
}

var FPSPEED = []string{
	"603.bwaves_s",
	"607.cactuBSSN_s",
	"619.lbm_s",
	"621.wrf_s",
	"627.cam4_s",
	"628.pop2_s",
	"638.imagick_s",
	"644.nab_s",
	"649.fotonik3d_s",
	"654.roms_s",
}

var ALLINT = append(INTRATE, INTSPEED...)
var ALLFP = append(FPRATE, FPSPEED...)
var ALL = append(ALLINT, ALLFP...)

var Benchmarks2Suite = map[string]string{
	"500.perlbench_r": "intrate",
	"502.gcc_r":       "intrate",
	"505.mcf_r":       "intrate",
	"520.omnetpp_r":   "intrate",
	"523.xalancbmk_r": "intrate",
	"525.x264_r":      "intrate",
	"531.deepsjeng_r": "intrate",
	"541.leela_r":     "intrate",
	"548.exchange2_r": "intrate",
	"557.xz_r":        "intrate",
	"600.perlbench_s": "intspeed",
	"602.gcc_s":       "intspeed",
	"605.mcf_s":       "intspeed",
	"620.omnetpp_s":   "intspeed",
	"623.xalancbmk_s": "intspeed",
	"625.x264_s":      "intspeed",
	"631.deepsjeng_s": "intspeed",
	"641.leela_s":     "intspeed",
	"648.exchange2_s": "intspeed",
	"657.xz_s":        "intspeed",
	"503.bwaves_r":    "fprate",
	"507.cactuBSSN_r": "fprate",
	"508.namd_r":      "fprate",
	"510.parest_r":    "fprate",
	"511.povray_r":    "fprate",
	"519.lbm_r":       "fprate",
	"521.wrf_r":       "fprate",
	"526.blender_r":   "fprate",
	"527.cam4_r":      "fprate",
	"538.imagick_r":   "fprate",
	"544.nab_r":       "fprate",
	"549.fotonik3d_r": "fprate",
	"554.roms_r":      "fprate",
	"603.bwaves_s":    "fpspeed",
	"607.cactuBSSN_s": "fpspeed",
	"619.lbm_s":       "fpspeed",
	"621.wrf_s":       "fpspeed",
	"627.cam4_s":      "fpspeed",
	"628.pop2_s":      "fpspeed",
	"638.imagick_s":   "fpspeed",
	"644.nab_s":       "fpspeed",
	"649.fotonik3d_s": "fpspeed",
	"654.roms_s":      "fpspeed",
}

func IsBenchmark(benchmark string) bool {
	for _, v := range ALL {
		if v == benchmark {
			return true
		}
	}
	return false
}
