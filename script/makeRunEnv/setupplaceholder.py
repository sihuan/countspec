#!/bin/python
import os

# 500.perlbench_r 502.gcc_r 503.bwaves_r 505.mcf_r 507.cactuBSSN_r 508.namd_r 510.parest_r 511.povray_r 519.lbm_r 520.omnetpp_r 521.wrf_r 523.xalancbmk_r 525.x264_r 526.blender_r 527.cam4_r 531.deepsjeng_r 538.imagick_r 541.leela_r 544.nab_r 548.exchange2_r 549.fotonik3d_r 554.roms_r 557.xz_r 600.perlbench_s 602.gcc_s 603.bwaves_s 605.mcf_s 607.cactuBSSN_s 619.lbm_s 620.omnetpp_s 621.wrf_s 623.xalancbmk_s 625.x264_s 627.cam4_s 628.pop2_s 631.deepsjeng_s 638.imagick_s 641.leela_s 644.nab_s 648.exchange2_s 649.fotonik3d_s 654.roms_s 657.xz_s
# 5xx - rate
ratedir = ["500.perlbench_r", "502.gcc_r", "503.bwaves_r", "505.mcf_r", "507.cactuBSSN_r", "508.namd_r", "510.parest_r", "511.povray_r", "519.lbm_r", "520.omnetpp_r", "521.wrf_r", "523.xalancbmk_r", "525.x264_r", "526.blender_r", "527.cam4_r", "531.deepsjeng_r", "538.imagick_r", "541.leela_r", "544.nab_r", "548.exchange2_r", "549.fotonik3d_r", "554.roms_r", "557.xz_r"]
# 6xx - speed
speeddir = ["600.perlbench_s", "602.gcc_s", "603.bwaves_s", "605.mcf_s", "607.cactuBSSN_s", "619.lbm_s", "620.omnetpp_s", "621.wrf_s", "623.xalancbmk_s", "625.x264_s", "627.cam4_s", "628.pop2_s", "631.deepsjeng_s", "638.imagick_s", "641.leela_s", "644.nab_s", "648.exchange2_s", "649.fotonik3d_s", "654.roms_s", "657.xz_s"]

label = "llvm-c9a6e993f7b3-rv64gc_zba_zbb_zbs-64"

def refdir(label, name):
    if name in ratedir:
        return f"run_base_refrate_{label}.0000"
    elif name in speeddir:
        return f"run_base_refspeed_{label}.0000"
    else:
        raise Exception(f"Unknown benchmark {name}")
    
def testdir(label, name):
    if name in ratedir + speeddir:
        return f"run_base_test_{label}.0000"
    else:
        raise Exception(f"Unknown benchmark {name}")

for name in ratedir + speeddir:
    ref = refdir(label, name)
    # 寻找 ref 文件夹里所有 名为 *.{label} 的可执行文件，删除并留下一个名为 *.mygo 的占位文件
    # 例如 500.perlbench_r/run_base_refrate_llvm-c9a6e993f7b3-rv64gc_zba_zbb_zbs-64.0000/perlbench_r_base.llvm-c9a6e993f7b3-rv64gc_zba_zbb_zbs-64 删除后，新建一个 perlbench_r_base.mygo 占位文件
    exe_files = [f for f in os.listdir(f"{name}/{ref}") if f.endswith(label)]
    for exe in exe_files:
        os.remove(f"{name}/{ref}/{exe}")
        with open(f"{name}/{ref}/{exe}.mygo", "w") as f:
            pass
    # 编辑 ref 文件夹下的 mygocmd 文件，把 其中的 ref 文件夹名改为 "ref"，然后再把其中的 label 改为 "mygo"
    content = ""
    with open(f"{name}/{ref}/{name}.mygocmd", "r") as f:
        content = f.read()
        content = content.replace(ref, "ref").replace(label, "mygo")
    with open(f"{name}/{ref}/{name}.mygocmd", "w") as f:
        f.write(content)

    # 重命名 ref 文件夹为 "ref" 文件夹
    os.rename(f"{name}/{ref}", f"{name}/ref")
    
        