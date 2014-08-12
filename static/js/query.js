$(function () {
                    $('#datetimepicker').datetimepicker({
                        pickTime: false,
                        language: 'zh-cn'
                    });
                    $('#datetimepicker2').datetimepicker({
                        pickTime: false,
                        language: 'zh-cn'
                    });
                    $('#table').dataTable({
                        "lengthMenu": [ [20, 25, 50, -1], [20, 25, 50, "全部"] ],
                        "pageLength": 20,
                        language: {
                                processing:     "处理中...",
                                search:         "查找",
                                lengthMenu:    "每页显示 _MENU_ 条数据",
                                info:           "从第 _START_ 条到第_END_条，共 _TOTAL_ 条记录",
                                infoEmpty:      "没有数据",
                                infoFiltered:   "(从_MAX_条数据中检索)",
                                infoPostFix:    "",
                                loadingRecords: "正在加载数据...",
                                zeroRecords:    "抱歉，没有找到",
                                emptyTable:     "表中没有可用的数据",
                                paginate: {
                                    first:      "第一页",
                                    previous:   "前一页",
                                    next:       "后一页",
                                    last:       "最后"
                                },
                                aria: {
                                    sortAscending:  "升序",
                                    sortDescending: "降序"
                                }
                        }, //end language
                        "footerCallback": function ( row, data, start, end, display ) {
                            var sum_gross = 0.0;
                            var sum_tare = 0.0;
                            var sum_suttle = 0.0;
                            var sum_tare_count = 0;
                            var sum_gross_count = 0;
                            for (var i = 0; i < data.length; i++) {
                                sum_gross += data[i][1]*1;
                                sum_tare += data[i][2]*1;
                                sum_suttle += data[i][3]*1;
                                sum_gross_count += data[i][4]*1;
                                sum_tare_count += data[i][5]*1;
                            };
                            $('#sum_gross').text(sum_gross.toFixed(2));
                            $('#sum_tare').text(sum_tare.toFixed(2));
                            $('#sum_suttle').text(sum_suttle.toFixed(2));
                            $('#sum_gross_count').text(sum_gross_count);
                            $('#sum_tare_count').text(sum_tare_count);
                        }, //end footerCallback
                    });//end dataTable
                });